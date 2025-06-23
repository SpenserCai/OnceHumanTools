<template>
  <div class="not-found-page">
    <div class="glitch-container">
      <h1 class="glitch-text" data-text="404">404</h1>
      <p class="error-message">页面未找到</p>
      <p class="error-desc">您访问的页面已经损坏或不存在</p>
      <router-link to="/" class="sci-fi-btn">
        返回首页
      </router-link>
    </div>
  </div>
</template>

<style lang="scss" scoped>
@use '@/styles/variables' as *;

.not-found-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: radial-gradient(ellipse at center, rgba(0, 255, 136, 0.1) 0%, transparent 70%);
}

.glitch-container {
  text-align: center;
  
  .glitch-text {
    font-family: $font-tech;
    font-size: 8rem;
    font-weight: 900;
    color: $primary-color;
    text-shadow: 
      0 0 10px currentColor,
      0 0 20px currentColor,
      0 0 40px currentColor;
    position: relative;
    margin-bottom: $spacing-md;
    
    &::before,
    &::after {
      content: attr(data-text);
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
    }
    
    &::before {
      animation: glitch-1 0.5s infinite;
      color: $secondary-color;
      z-index: -1;
    }
    
    &::after {
      animation: glitch-2 0.5s infinite;
      color: $accent-color;
      z-index: -2;
    }
  }
  
  .error-message {
    font-family: $font-tech;
    font-size: 2rem;
    color: $text-primary;
    margin-bottom: $spacing-sm;
    text-transform: uppercase;
    letter-spacing: 2px;
  }
  
  .error-desc {
    color: $text-secondary;
    margin-bottom: $spacing-xl;
  }
}

@keyframes glitch-1 {
  0%, 100% {
    clip: rect(0, 900px, 0, 0);
    transform: translate(0);
  }
  20% {
    clip: rect(20px, 900px, 60px, 0);
    transform: translate(-2px, 2px);
  }
  40% {
    clip: rect(65px, 900px, 120px, 0);
    transform: translate(2px, -2px);
  }
  60% {
    clip: rect(120px, 900px, 180px, 0);
    transform: translate(-2px, 2px);
  }
  80% {
    clip: rect(25px, 900px, 80px, 0);
    transform: translate(2px, -2px);
  }
}

@keyframes glitch-2 {
  0%, 100% {
    clip: rect(0, 900px, 0, 0);
    transform: translate(0);
  }
  20% {
    clip: rect(35px, 900px, 90px, 0);
    transform: translate(2px, -2px);
  }
  40% {
    clip: rect(85px, 900px, 140px, 0);
    transform: translate(-2px, 2px);
  }
  60% {
    clip: rect(140px, 900px, 200px, 0);
    transform: translate(2px, -2px);
  }
  80% {
    clip: rect(45px, 900px, 100px, 0);
    transform: translate(-2px, 2px);
  }
}
</style>