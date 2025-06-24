<template>
  <div class="home-page">
    <!-- 顶部导航 -->
    <nav class="sci-fi-nav">
      <div class="nav-container">
        <div class="logo">
          <span class="logo-text glow-text">ONCEHUMAN</span>
          <span class="logo-sub">TOOLS</span>
        </div>
        <div class="nav-links">
          <router-link to="/" class="nav-link active">首页</router-link>
          <router-link to="/tools" class="nav-link">工具</router-link>
          <router-link to="/about" class="nav-link">关于</router-link>
        </div>
      </div>
    </nav>
    
    <!-- 主标题区域 -->
    <section class="hero-section">
      <div class="hero-content">
        <h1 class="hero-title fade-in">
          <span class="title-line">ONCEHUMAN</span>
          <span class="title-line accent">工具集</span>
        </h1>
        <p class="hero-subtitle fade-in">
          专业的游戏数据计算与分析工具
        </p>
        <div class="hero-buttons fade-in">
          <HologramButton variant="primary" @click="$router.push('/tools')">
            开始使用
          </HologramButton>
          <HologramButton variant="outline" @click="scrollToFeatures">
            了解更多
          </HologramButton>
        </div>
      </div>
      
      <!-- 3D背景装饰 -->
      <div class="hero-bg">
        <canvas ref="bgCanvas"></canvas>
      </div>
    </section>
    
    <!-- 功能展示 -->
    <section id="features" class="features-section">
      <div class="container">
        <h2 class="section-title glow-text">核心功能</h2>
        <div class="features-grid">
          <HologramCard
            v-for="feature in features" 
            :key="feature.id"
            class="feature-card"
            variant="primary"
            interactive
            @click="navigateToTool(feature.path)"
          >
            <div class="feature-icon">
              <component :is="feature.icon" />
            </div>
            <h3 class="feature-title">{{ feature.title }}</h3>
            <p class="feature-desc">{{ feature.description }}</p>
          </HologramCard>
        </div>
      </div>
    </section>
    
    <!-- 统计数据 -->
    <section class="stats-section grid-bg">
      <div class="container">
        <div class="stats-grid">
          <div v-for="stat in stats" :key="stat.label" class="stat-item">
            <div class="stat-value glow-text">{{ stat.value }}</div>
            <div class="stat-label">{{ stat.label }}</div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { Histogram, TrendCharts, DataAnalysis, Setting } from '@element-plus/icons-vue'
import * as THREE from 'three'
import { HologramButton, HologramCard } from '@/components'

const router = useRouter()
const bgCanvas = ref(null)
let animationId = null
let renderer, scene, camera

const features = [
  {
    id: 'affix-probability',
    icon: Histogram,
    title: '词条概率计算器',
    description: '精确计算模组词条组合出现的概率，优化你的装备选择策略',
    path: '/tools/affix-probability'
  },
  {
    id: 'strengthen-probability',
    icon: TrendCharts,
    title: '强化概率计算器',
    description: '计算模组强化到目标等级的成功率，规划最优强化路径',
    path: '/tools/strengthen-probability'
  },
  {
    id: 'more-tools',
    icon: DataAnalysis,
    title: '更多工具',
    description: '更多实用工具正在开发中，敬请期待',
    path: '/tools'
  }
]

const stats = ref([
  { label: '计算次数', value: '10,234' },
  { label: '活跃用户', value: '1,528' },
  { label: '工具数量', value: '2+' },
  { label: '准确率', value: '99.9%' }
])

const navigateToTool = (path) => {
  router.push(path)
}

const scrollToFeatures = () => {
  document.getElementById('features')?.scrollIntoView({ behavior: 'smooth' })
}

// 初始化3D背景
const init3DBackground = () => {
  if (!bgCanvas.value) return
  
  // 创建场景
  scene = new THREE.Scene()
  scene.fog = new THREE.Fog(0x000000, 1, 1000)
  
  // 创建相机
  camera = new THREE.PerspectiveCamera(
    75,
    window.innerWidth / window.innerHeight,
    0.1,
    1000
  )
  camera.position.z = 30
  
  // 创建渲染器
  renderer = new THREE.WebGLRenderer({
    canvas: bgCanvas.value,
    alpha: true,
    antialias: true
  })
  renderer.setSize(window.innerWidth, window.innerHeight)
  renderer.setPixelRatio(window.devicePixelRatio)
  
  // 创建粒子系统
  const geometry = new THREE.BufferGeometry()
  const vertices = []
  
  for (let i = 0; i < 5000; i++) {
    vertices.push(
      THREE.MathUtils.randFloatSpread(200),
      THREE.MathUtils.randFloatSpread(200),
      THREE.MathUtils.randFloatSpread(200)
    )
  }
  
  geometry.setAttribute('position', new THREE.Float32BufferAttribute(vertices, 3))
  
  const material = new THREE.PointsMaterial({
    color: 0x00d4ff,
    size: 0.5,
    transparent: true,
    opacity: 0.8,
    blending: THREE.AdditiveBlending
  })
  
  const particles = new THREE.Points(geometry, material)
  scene.add(particles)
  
  // 动画循环
  const animate = () => {
    animationId = requestAnimationFrame(animate)
    
    particles.rotation.x += 0.0005
    particles.rotation.y += 0.001
    
    renderer.render(scene, camera)
  }
  
  animate()
}

// 处理窗口大小变化
const handleResize = () => {
  if (camera && renderer) {
    camera.aspect = window.innerWidth / window.innerHeight
    camera.updateProjectionMatrix()
    renderer.setSize(window.innerWidth, window.innerHeight)
  }
}

onMounted(() => {
  init3DBackground()
  window.addEventListener('resize', handleResize)
})

onUnmounted(() => {
  if (animationId) {
    cancelAnimationFrame(animationId)
  }
  if (renderer) {
    renderer.dispose()
  }
  window.removeEventListener('resize', handleResize)
})
</script>

<style lang="scss" scoped>
@use '@/styles/variables' as *;

.home-page {
  min-height: 100vh;
}

// 导航栏
.sci-fi-nav {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: $z-header;
  background: rgba(0, 0, 0, 0.8);
  backdrop-filter: blur(10px);
  border-bottom: 1px solid $border-color;
  
  .nav-container {
    max-width: 1200px;
    margin: 0 auto;
    padding: $spacing-md $spacing-lg;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  
  .logo {
    display: flex;
    align-items: baseline;
    gap: $spacing-sm;
    
    .logo-text {
      font-family: $font-tech;
      font-size: 1.5rem;
      font-weight: 900;
      letter-spacing: 2px;
    }
    
    .logo-sub {
      font-family: $font-tech;
      font-size: 0.8rem;
      color: $text-muted;
      letter-spacing: 1px;
    }
  }
  
  .nav-links {
    display: flex;
    gap: $spacing-xl;
    
    .nav-link {
      color: $text-secondary;
      text-decoration: none;
      font-family: $font-tech;
      text-transform: uppercase;
      letter-spacing: 1px;
      transition: color $transition-fast;
      
      &:hover,
      &.active {
        color: $primary-color;
      }
    }
  }
}

// 主标题区域
.hero-section {
  position: relative;
  height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  
  .hero-content {
    position: relative;
    z-index: 2;
    text-align: center;
    padding: $spacing-xl;
  }
  
  .hero-title {
    font-family: $font-tech;
    font-size: clamp(3rem, 10vw, 6rem);
    font-weight: 900;
    line-height: 1.1;
    margin-bottom: $spacing-lg;
    
    .title-line {
      display: block;
      
      &.accent {
        color: $primary-color;
        text-shadow: 0 0 30px currentColor;
      }
    }
  }
  
  .hero-subtitle {
    font-size: 1.2rem;
    color: $text-secondary;
    margin-bottom: $spacing-xxl;
    opacity: 0.8;
  }
  
  .hero-buttons {
    display: flex;
    gap: $spacing-lg;
    justify-content: center;
    flex-wrap: wrap;
  }
  
  .hero-bg {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 1;
    opacity: 0.5;
    
    canvas {
      width: 100%;
      height: 100%;
    }
  }
}

// 功能展示
.features-section {
  padding: $spacing-xxl * 2 $spacing-lg;
  background: $bg-secondary;
  
  .container {
    max-width: 1200px;
    margin: 0 auto;
  }
  
  .section-title {
    font-family: $font-tech;
    font-size: 2.5rem;
    text-align: center;
    margin-bottom: $spacing-xxl;
    text-transform: uppercase;
    letter-spacing: 2px;
  }
  
  .features-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: $spacing-xl;
  }
  
  .feature-card {
    cursor: pointer;
    text-align: center;
    
    .feature-icon {
      font-size: 3rem;
      color: $primary-color;
      margin-bottom: $spacing-lg;
    }
    
    .feature-title {
      font-family: $font-tech;
      font-size: 1.3rem;
      margin-bottom: $spacing-md;
      color: $text-primary;
    }
    
    .feature-desc {
      color: $text-secondary;
      line-height: 1.6;
    }
  }
}

// 统计数据
.stats-section {
  padding: $spacing-xxl * 2 $spacing-lg;
  background: $bg-primary;
  
  .container {
    max-width: 1200px;
    margin: 0 auto;
  }
  
  .stats-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
    gap: $spacing-xl;
    text-align: center;
  }
  
  .stat-item {
    .stat-value {
      font-family: $font-tech;
      font-size: 2.5rem;
      font-weight: 700;
      margin-bottom: $spacing-sm;
    }
    
    .stat-label {
      color: $text-secondary;
      text-transform: uppercase;
      letter-spacing: 1px;
    }
  }
}

// 响应式
@media (max-width: 768px) {
  .sci-fi-nav {
    .nav-container {
      flex-direction: column;
      gap: $spacing-md;
    }
  }
  
  .hero-section {
    .hero-buttons {
      flex-direction: column;
      align-items: center;
    }
  }
}
</style>