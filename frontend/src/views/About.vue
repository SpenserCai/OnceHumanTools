<template>
  <div class="about-page">
    <div class="container">
      <h1 class="page-title glow-text">关于 OnceHuman 工具集</h1>
      
      <HologramCard class="content-section" title="项目介绍" variant="primary">
        <p>
          OnceHuman工具集是一个专为《七日世界》(Once Human)游戏玩家打造的综合性工具平台。
          我们提供精确的数学计算工具，帮助玩家优化游戏策略，提升游戏体验。
        </p>
      </HologramCard>
      
      <HologramCard class="content-section" title="功能特色" variant="secondary">
        <ul>
          <li>🎯 精确的概率计算 - 基于游戏机制的数学模型</li>
          <li>🚀 高性能后端 - 使用Go语言开发，响应迅速</li>
          <li>🎨 炫酷UI设计 - 沉浸式科幻风格界面</li>
          <li>🤖 Discord集成 - 在聊天中即可使用工具</li>
          <li>📱 响应式设计 - 支持各种设备访问</li>
        </ul>
      </HologramCard>
      
      <HologramCard class="content-section" title="技术栈" variant="primary">
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
      </HologramCard>
      
      <HologramCard class="content-section" title="联系我们" variant="secondary">
        <p>
          如果您有任何问题、建议或想要贡献代码，欢迎通过以下方式联系我们：
        </p>
        <div class="contact-links">
          <HologramButton variant="outline" @click="openGitHub">
            <Link /> GitHub
          </HologramButton>
          <HologramButton variant="outline" disabled>
            <Message /> Discord
          </HologramButton>
        </div>
      </HologramCard>
    </div>
  </div>
</template>

<script setup>
import { Link, Message } from '@element-plus/icons-vue'
import { HologramCard, HologramButton } from '@/components'

const openGitHub = () => {
  window.open('https://github.com/SpenserCai/OnceHumanTools', '_blank')
}
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
  }
}
</style>