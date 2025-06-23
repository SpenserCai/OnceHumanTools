<template>
  <div id="app" class="app-container">
    <!-- 背景粒子效果 -->
    <particles
      id="tsparticles"
      class="particles-bg"
      :options="particlesOptions"
    />
    
    <!-- 主内容区域 -->
    <div class="main-content">
      <transition name="fade" mode="out-in">
        <router-view />
      </transition>
    </div>
    
    <!-- 科幻风格装饰元素 -->
    <div class="sci-fi-border top-left"></div>
    <div class="sci-fi-border top-right"></div>
    <div class="sci-fi-border bottom-left"></div>
    <div class="sci-fi-border bottom-right"></div>
  </div>
</template>

<script setup>
const particlesOptions = {
  background: {
    color: {
      value: '#000000'
    }
  },
  fpsLimit: 120,
  particles: {
    color: {
      value: '#00ff88'
    },
    links: {
      color: '#00ff88',
      distance: 150,
      enable: true,
      opacity: 0.2,
      width: 1
    },
    move: {
      enable: true,
      speed: 1,
      direction: 'none',
      random: false,
      straight: false,
      outModes: {
        default: 'bounce'
      }
    },
    number: {
      density: {
        enable: true,
        area: 800
      },
      value: 80
    },
    opacity: {
      value: 0.5
    },
    shape: {
      type: 'circle'
    },
    size: {
      value: { min: 1, max: 3 }
    }
  },
  detectRetina: true
}
</script>

<style lang="scss">
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: 'Noto Sans SC', sans-serif;
  background: #000;
  color: #fff;
  overflow-x: hidden;
}

.app-container {
  position: relative;
  min-height: 100vh;
  width: 100%;
}

.particles-bg {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 0;
}

.main-content {
  position: relative;
  z-index: 10;
  min-height: 100vh;
}

// 科幻风格边框装饰
.sci-fi-border {
  position: fixed;
  width: 100px;
  height: 100px;
  border: 2px solid #00ff88;
  z-index: 20;
  pointer-events: none;
  
  &::before,
  &::after {
    content: '';
    position: absolute;
    background: #00ff88;
  }
  
  &.top-left {
    top: 20px;
    left: 20px;
    border-right: none;
    border-bottom: none;
    
    &::before {
      top: -2px;
      left: -2px;
      width: 20px;
      height: 2px;
    }
    
    &::after {
      top: -2px;
      left: -2px;
      width: 2px;
      height: 20px;
    }
  }
  
  &.top-right {
    top: 20px;
    right: 20px;
    border-left: none;
    border-bottom: none;
    
    &::before {
      top: -2px;
      right: -2px;
      width: 20px;
      height: 2px;
    }
    
    &::after {
      top: -2px;
      right: -2px;
      width: 2px;
      height: 20px;
    }
  }
  
  &.bottom-left {
    bottom: 20px;
    left: 20px;
    border-right: none;
    border-top: none;
    
    &::before {
      bottom: -2px;
      left: -2px;
      width: 20px;
      height: 2px;
    }
    
    &::after {
      bottom: -2px;
      left: -2px;
      width: 2px;
      height: 20px;
    }
  }
  
  &.bottom-right {
    bottom: 20px;
    right: 20px;
    border-left: none;
    border-top: none;
    
    &::before {
      bottom: -2px;
      right: -2px;
      width: 20px;
      height: 2px;
    }
    
    &::after {
      bottom: -2px;
      right: -2px;
      width: 2px;
      height: 20px;
    }
  }
}

// 页面过渡动画
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.3s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

// 滚动条样式
::-webkit-scrollbar {
  width: 8px;
  height: 8px;
}

::-webkit-scrollbar-track {
  background: rgba(0, 0, 0, 0.5);
}

::-webkit-scrollbar-thumb {
  background: #00ff88;
  border-radius: 4px;
  
  &:hover {
    background: #00cc66;
  }
}
</style>