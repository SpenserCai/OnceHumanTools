<template>
  <div class="tools-page">
    <router-view v-if="$route.matched.length > 1" />
    <div v-else class="tools-list">
      <h1 class="page-title glow-text">工具列表</h1>
      <div class="tools-grid">
        <router-link
          v-for="tool in tools"
          :key="tool.path"
          :to="tool.path"
          class="tool-card sci-fi-card"
        >
          <div class="tool-icon">
            <component :is="tool.icon" />
          </div>
          <h3 class="tool-title">{{ tool.name }}</h3>
          <p class="tool-desc">{{ tool.description }}</p>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { Histogram, TrendCharts } from '@element-plus/icons-vue'

const tools = [
  {
    name: '词条概率计算器',
    description: '计算特定词条组合出现的概率',
    path: '/tools/affix-probability',
    icon: Histogram
  },
  {
    name: '强化概率计算器',
    description: '计算模组强化到目标等级的概率',
    path: '/tools/strengthen-probability',
    icon: TrendCharts
  }
]
</script>

<style lang="scss" scoped>
@use '@/styles/variables' as *;

.tools-page {
  min-height: 100vh;
  padding: 80px $spacing-lg;
}

.tools-list {
  max-width: 1200px;
  margin: 0 auto;
  
  .page-title {
    font-family: $font-tech;
    font-size: 2.5rem;
    text-align: center;
    margin-bottom: $spacing-xxl;
    text-transform: uppercase;
    letter-spacing: 2px;
  }
  
  .tools-grid {
    display: grid;
    grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
    gap: $spacing-xl;
  }
  
  .tool-card {
    display: block;
    text-decoration: none;
    text-align: center;
    transition: transform $transition-normal;
    
    &:hover {
      transform: translateY(-5px);
    }
    
    .tool-icon {
      font-size: 3rem;
      color: $primary-color;
      margin-bottom: $spacing-md;
    }
    
    .tool-title {
      font-family: $font-tech;
      font-size: 1.3rem;
      color: $text-primary;
      margin-bottom: $spacing-sm;
    }
    
    .tool-desc {
      color: $text-secondary;
    }
  }
}
</style>