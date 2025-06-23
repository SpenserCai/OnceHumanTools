<template>
  <div class="about-page">
    <div class="container">
      <h1 class="page-title glow-text">关于 OnceHuman 工具集</h1>
      
      <div class="content-section sci-fi-card">
        <h2>项目介绍</h2>
        <p>
          OnceHuman工具集是一个专为《七日世界》(Once Human)游戏玩家打造的综合性工具平台。
          我们提供精确的数学计算工具，帮助玩家优化游戏策略，提升游戏体验。
        </p>
      </div>
      
      <div class="content-section sci-fi-card">
        <h2>功能特色</h2>
        <ul>
          <li>🎯 精确的概率计算 - 基于游戏机制的数学模型</li>
          <li>🚀 高性能后端 - 使用Go语言开发，响应迅速</li>
          <li>🎨 炫酷UI设计 - 沉浸式科幻风格界面</li>
          <li>🤖 Discord集成 - 在聊天中即可使用工具</li>
          <li>📱 响应式设计 - 支持各种设备访问</li>
        </ul>
      </div>
      
      <div class="content-section sci-fi-card">
        <h2>技术栈</h2>
        <div class="tech-grid">
          <div class="tech-item">
            <h3>前端</h3>
            <p>Vue 3 + Vite + Element Plus</p>
          </div>
          <div class="tech-item">
            <h3>后端</h3>
            <p>Go + go-swagger + RESTful API</p>
          </div>
          <div class="tech-item">
            <h3>机器人</h3>
            <p>Discord.go + 命令模式</p>
          </div>
        </div>
      </div>
      
      <div class="content-section sci-fi-card">
        <h2>联系我们</h2>
        <p>
          如果您有任何问题、建议或想要贡献代码，欢迎通过以下方式联系我们：
        </p>
        <div class="contact-links">
          <a href="https://github.com/oncehuman/tools" target="_blank" class="contact-link">
            <Github /> GitHub
          </a>
          <a href="#" class="contact-link">
            <Message /> Discord
          </a>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Github, Message } from '@element-plus/icons-vue'
</script>

<style lang="scss" scoped>
@use '@/styles/variables' as *;

.about-page {
  min-height: 100vh;
  padding: 80px $spacing-lg $spacing-xxl;
  
  .container {
    max-width: 800px;
    margin: 0 auto;
  }
  
  .page-title {
    font-family: $font-tech;
    font-size: 2.5rem;
    text-align: center;
    margin-bottom: $spacing-xxl;
    text-transform: uppercase;
    letter-spacing: 2px;
  }
  
  .content-section {
    margin-bottom: $spacing-xl;
    
    h2 {
      font-family: $font-tech;
      font-size: 1.5rem;
      color: $primary-color;
      margin-bottom: $spacing-md;
    }
    
    p {
      line-height: 1.8;
      color: $text-secondary;
      margin-bottom: $spacing-md;
    }
    
    ul {
      list-style: none;
      padding: 0;
      
      li {
        padding: $spacing-sm 0;
        color: $text-secondary;
        
        &::before {
          content: '▸';
          color: $primary-color;
          margin-right: $spacing-sm;
        }
      }
    }
  }
  
  .tech-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    gap: $spacing-lg;
    margin-top: $spacing-lg;
    
    .tech-item {
      text-align: center;
      
      h3 {
        color: $primary-color;
        margin-bottom: $spacing-sm;
      }
      
      p {
        color: $text-secondary;
      }
    }
  }
  
  .contact-links {
    display: flex;
    gap: $spacing-lg;
    margin-top: $spacing-lg;
    
    .contact-link {
      display: flex;
      align-items: center;
      gap: $spacing-sm;
      padding: $spacing-sm $spacing-lg;
      border: 1px solid $border-color;
      color: $text-primary;
      text-decoration: none;
      transition: all $transition-normal;
      
      &:hover {
        border-color: $primary-color;
        color: $primary-color;
        box-shadow: $shadow-sm;
      }
    }
  }
}
</style>