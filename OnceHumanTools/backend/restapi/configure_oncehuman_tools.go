// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/rs/cors"

	"github.com/oncehuman/tools/internal/handlers"
	"github.com/oncehuman/tools/restapi/operations"
	"github.com/oncehuman/tools/restapi/operations/mod"
	"github.com/oncehuman/tools/restapi/operations/system"
	"github.com/oncehuman/tools/restapi/operations/tools"
)

//go:generate swagger generate server --target ../../backend --name OncehumanTools --spec ../api/swagger.yaml --principal interface{} --exclude-main

func configureFlags(api *operations.OncehumanToolsAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.OncehumanToolsAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	// 创建处理器实例
	modHandler := handlers.NewModHandler()
	systemHandler := handlers.NewSystemHandler()
	toolsHandler := handlers.NewToolsHandler()

	// 健康检查
	api.SystemHealthCheckHandler = system.HealthCheckHandlerFunc(func(params system.HealthCheckParams) middleware.Responder {
		return systemHandler.HealthCheck(params)
	})

	// 词条列表
	api.ModListAffixesHandler = mod.ListAffixesHandlerFunc(func(params mod.ListAffixesParams) middleware.Responder {
		return modHandler.ListAffixes(params)
	})

	// 词条概率计算
	api.ModCalculateAffixProbabilityHandler = mod.CalculateAffixProbabilityHandlerFunc(func(params mod.CalculateAffixProbabilityParams) middleware.Responder {
		return modHandler.CalculateAffixProbability(params)
	})

	// 强化概率计算
	api.ModCalculateStrengthenProbabilityHandler = mod.CalculateStrengthenProbabilityHandlerFunc(func(params mod.CalculateStrengthenProbabilityParams) middleware.Responder {
		return modHandler.CalculateStrengthenProbability(params)
	})

	// 工具列表
	api.ToolsListToolsHandler = tools.ListToolsHandlerFunc(func(params tools.ListToolsParams) middleware.Responder {
		return toolsHandler.ListTools(params)
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
	// 设置服务器超时
	s.ReadTimeout = 30 * time.Second
	s.WriteTimeout = 30 * time.Second
	s.IdleTimeout = 120 * time.Second
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	// 添加CORS支持
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	
	return c.Handler(handler)
}
