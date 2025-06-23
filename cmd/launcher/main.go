package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"
)

var (
	// 命令行参数
	mode           = flag.String("mode", "integrated", "运行模式: integrated 或 standalone")
	enableBot      = flag.Bool("bot", false, "是否启动Discord机器人")
	frontendPath   = flag.String("frontend", "./frontend", "前端文件路径")
	backendPort    = flag.String("backend-port", "8080", "后端服务端口")
	backendScheme  = flag.String("backend-scheme", "http", "后端服务协议: http 或 https")
	tlsCertificate = flag.String("tls-certificate", "", "TLS证书文件路径（HTTPS模式需要）")
	tlsKey         = flag.String("tls-key", "", "TLS私钥文件路径（HTTPS模式需要）")
	tlsCA          = flag.String("tls-ca", "", "TLS CA证书文件路径（可选）")
	frontendPort   = flag.String("frontend-port", "3000", "前端服务端口")
	launcherPort   = flag.String("port", "9000", "启动器端口（集成模式）")
)

type Launcher struct {
	ctx        context.Context
	cancel     context.CancelFunc
	wg         sync.WaitGroup
	backendCmd *exec.Cmd
	botCmd     *exec.Cmd
}

func NewLauncher() *Launcher {
	ctx, cancel := context.WithCancel(context.Background())
	return &Launcher{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (l *Launcher) Start() error {
	flag.Parse()

	log.Printf("启动模式: %s", *mode)

	switch *mode {
	case "integrated":
		return l.runIntegrated()
	case "standalone":
		return l.runStandalone()
	default:
		return fmt.Errorf("未知的运行模式: %s", *mode)
	}
}

// runIntegrated 集成模式：同时运行前端、后端和可选的机器人
func (l *Launcher) runIntegrated() error {
	// 启动后端服务
	if err := l.startBackend(); err != nil {
		return fmt.Errorf("启动后端失败: %w", err)
	}

	// 等待后端启动
	if err := l.waitForBackend(); err != nil {
		return fmt.Errorf("后端启动超时: %w", err)
	}

	// 启动Discord机器人（如果启用）
	if *enableBot {
		if err := l.startBot(); err != nil {
			log.Printf("启动机器人失败: %v", err)
		}
	}

	// 启动集成Web服务器
	go l.startIntegratedServer()

	log.Printf("OnceHuman工具集已启动在 http://localhost:%s", *launcherPort)
	log.Println("按 Ctrl+C 停止所有服务")

	// 等待中断信号
	l.waitForShutdown()

	return nil
}

// runStandalone 独立模式：只作为启动器，不运行服务
func (l *Launcher) runStandalone() error {
	log.Println("独立模式：请分别启动各个服务")
	log.Printf("前端路径: %s", *frontendPath)
	log.Printf("后端端口: %s", *backendPort)
	log.Printf("后端协议: %s", *backendScheme)

	if *backendScheme == "https" {
		log.Printf("TLS证书: %s", *tlsCertificate)
		log.Printf("TLS私钥: %s", *tlsKey)
		if *tlsCA != "" {
			log.Printf("TLS CA: %s", *tlsCA)
		}
	}

	if *enableBot {
		log.Println("Discord机器人: 已启用")
	}
	return nil
}

// startBackend 启动后端服务
func (l *Launcher) startBackend() error {
	log.Println("正在启动后端服务...")

	// 查找后端可执行文件
	backendPath := "./backend/server"
	if _, err := os.Stat(backendPath); os.IsNotExist(err) {
		// 尝试从release目录
		backendPath = "./release/backend/server"
		if _, err := os.Stat(backendPath); os.IsNotExist(err) {
			return fmt.Errorf("找不到后端可执行文件")
		}
	}

	// 构建命令参数
	args := []string{
		"--scheme=" + *backendScheme,
		"--host=0.0.0.0",
		"--port", *backendPort,
	}

	// 如果使用 HTTPS，验证并添加 TLS 参数
	if *backendScheme == "https" {
		if *tlsCertificate == "" || *tlsKey == "" {
			return fmt.Errorf("HTTPS模式需要提供 --tls-certificate 和 --tls-key 参数")
		}

		// 验证证书文件是否存在
		if _, err := os.Stat(*tlsCertificate); os.IsNotExist(err) {
			return fmt.Errorf("TLS证书文件不存在: %s", *tlsCertificate)
		}
		if _, err := os.Stat(*tlsKey); os.IsNotExist(err) {
			return fmt.Errorf("TLS私钥文件不存在: %s", *tlsKey)
		}

		args = append(args, "--tls-certificate", *tlsCertificate)
		args = append(args, "--tls-key", *tlsKey)

		// 可选的 CA 证书
		if *tlsCA != "" {
			if _, err := os.Stat(*tlsCA); os.IsNotExist(err) {
				return fmt.Errorf("TLS CA证书文件不存在: %s", *tlsCA)
			}
			args = append(args, "--tls-ca", *tlsCA)
		}

		log.Printf("使用 HTTPS 模式启动，证书: %s", *tlsCertificate)
	} else {
		log.Println("使用 HTTP 模式启动（开发环境）")
	}

	l.backendCmd = exec.CommandContext(l.ctx, backendPath, args...)
	l.backendCmd.Stdout = os.Stdout
	l.backendCmd.Stderr = os.Stderr
	l.backendCmd.Env = append(os.Environ(), fmt.Sprintf("PORT=%s", *backendPort))

	if err := l.backendCmd.Start(); err != nil {
		return fmt.Errorf("启动后端进程失败: %w", err)
	}

	l.wg.Add(1)
	go func() {
		defer l.wg.Done()
		if err := l.backendCmd.Wait(); err != nil {
			log.Printf("后端进程退出: %v", err)
		}
	}()

	return nil
}

// startBot 启动Discord机器人
func (l *Launcher) startBot() error {
	log.Println("正在启动Discord机器人...")

	// 查找机器人可执行文件
	botPath := "./bot/bot"
	if _, err := os.Stat(botPath); os.IsNotExist(err) {
		// 尝试从release目录
		botPath = "./release/bot/bot"
		if _, err := os.Stat(botPath); os.IsNotExist(err) {
			return fmt.Errorf("找不到机器人可执行文件")
		}
	}

	l.botCmd = exec.CommandContext(l.ctx, botPath)
	l.botCmd.Stdout = os.Stdout
	l.botCmd.Stderr = os.Stderr
	l.botCmd.Env = os.Environ()

	if err := l.botCmd.Start(); err != nil {
		return fmt.Errorf("启动机器人进程失败: %w", err)
	}

	l.wg.Add(1)
	go func() {
		defer l.wg.Done()
		if err := l.botCmd.Wait(); err != nil {
			log.Printf("机器人进程退出: %v", err)
		}
	}()

	return nil
}

// waitForBackend 等待后端服务启动
func (l *Launcher) waitForBackend() error {
	backendURL := fmt.Sprintf("%s://localhost:%s/api/v1/health", *backendScheme, *backendPort)

	log.Printf("等待后端服务启动: %s", backendURL)

	for i := 0; i < 30; i++ {
		resp, err := http.Get(backendURL)
		if err == nil && resp.StatusCode == http.StatusOK {
			resp.Body.Close()
			log.Println("后端服务已就绪")
			return nil
		}
		if resp != nil {
			resp.Body.Close()
		}
		time.Sleep(1 * time.Second)
	}

	return fmt.Errorf("后端服务启动超时")
}

// startIntegratedServer 启动集成Web服务器
func (l *Launcher) startIntegratedServer() {
	mux := http.NewServeMux()

	// 配置API代理
	apiURL, _ := url.Parse(fmt.Sprintf("%s://localhost:%s", *backendScheme, *backendPort))
	apiProxy := httputil.NewSingleHostReverseProxy(apiURL)

	// API路由
	mux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		apiProxy.ServeHTTP(w, r)
	})

	// Swagger文档路由
	mux.HandleFunc("/docs", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, fmt.Sprintf("%s://localhost:%s/docs", *backendScheme, *backendPort), http.StatusTemporaryRedirect)
	})

	// 静态文件服务
	staticPath, _ := filepath.Abs(*frontendPath)
	fs := http.FileServer(http.Dir(staticPath))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 检查文件是否存在
		path := filepath.Join(staticPath, r.URL.Path)
		_, err := os.Stat(path)

		// 如果文件不存在且不是API请求，返回index.html（支持Vue Router）
		if os.IsNotExist(err) && r.URL.Path != "/" {
			http.ServeFile(w, r, filepath.Join(staticPath, "index.html"))
			return
		}

		fs.ServeHTTP(w, r)
	})

	server := &http.Server{
		Addr:    ":" + *launcherPort,
		Handler: mux,
	}

	log.Printf("集成服务器启动在端口 %s", *launcherPort)
	log.Printf("API代理: %s://localhost:%s -> http://localhost:%s/api/", *backendScheme, *backendPort, *launcherPort)

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("集成服务器错误: %v", err)
	}
}

// waitForShutdown 等待关闭信号
func (l *Launcher) waitForShutdown() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	log.Println("收到关闭信号，正在停止服务...")

	// 取消上下文
	l.cancel()

	// 等待所有子进程退出
	done := make(chan struct{})
	go func() {
		l.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		log.Println("所有服务已停止")
	case <-time.After(10 * time.Second):
		log.Println("强制停止服务")
		if l.backendCmd != nil && l.backendCmd.Process != nil {
			l.backendCmd.Process.Kill()
		}
		if l.botCmd != nil && l.botCmd.Process != nil {
			l.botCmd.Process.Kill()
		}
	}
}

func main() {
	launcher := NewLauncher()
	if err := launcher.Start(); err != nil {
		log.Fatal(err)
	}
}
