<template>
  <span 
    :class="[
      'hologram-tag',
      `hologram-tag--${type}`,
      `hologram-tag--${size}`,
      {
        'hologram-tag--glow': glow
      }
    ]"
  >
    <span class="tag-content">
      <slot></slot>
    </span>
    <div class="tag-glow"></div>
  </span>
</template>

<script setup>
const props = defineProps({
  type: {
    type: String,
    default: 'default',
    validator: (value) => ['default', 'success', 'warning', 'danger', 'info'].includes(value)
  },
  size: {
    type: String,
    default: 'medium',
    validator: (value) => ['small', 'medium', 'large'].includes(value)
  },
  glow: {
    type: Boolean,
    default: false
  }
})
</script>

<style lang="scss" scoped>
@use '@/styles/variables' as *;

.hologram-tag {
  position: relative;
  display: inline-flex;
  align-items: center;
  border-radius: $radius-sm;
  font-weight: 500;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  overflow: hidden;
  transition: all $transition-normal;
  
  .tag-content {
    position: relative;
    z-index: 2;
    padding: inherit;
  }
  
  .tag-glow {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    border-radius: inherit;
    opacity: 0;
    transition: opacity $transition-normal;
    pointer-events: none;
  }
  
  // Sizes
  &--small {
    padding: 2px $spacing-xs;
    font-size: 0.625rem;
  }
  
  &--medium {
    padding: $spacing-xs $spacing-sm;
    font-size: 0.75rem;
  }
  
  &--large {
    padding: $spacing-sm $spacing-md;
    font-size: 0.875rem;
  }
  
  // Types
  &--default {
    background: rgba(0, 212, 255, 0.2);
    color: #66ddff;
    border: 1px solid rgba(0, 212, 255, 0.4);
    
    .tag-glow {
      background: rgba(0, 212, 255, 0.3);
      box-shadow: 0 0 15px rgba(0, 212, 255, 0.4);
    }
  }
  
  &--success {
    background: rgba(0, 255, 170, 0.2);
    color: #00ffaa;
    border: 1px solid rgba(0, 255, 170, 0.4);
    
    .tag-glow {
      background: rgba(0, 255, 170, 0.3);
      box-shadow: 0 0 15px rgba(0, 255, 170, 0.4);
    }
  }
  
  &--warning {
    background: rgba(255, 170, 0, 0.2);
    color: #ffaa00;
    border: 1px solid rgba(255, 170, 0, 0.4);
    
    .tag-glow {
      background: rgba(255, 170, 0, 0.3);
      box-shadow: 0 0 15px rgba(255, 170, 0, 0.4);
    }
  }
  
  &--danger {
    background: rgba(255, 68, 68, 0.2);
    color: #ff6b6b;
    border: 1px solid rgba(255, 107, 107, 0.4);
    
    .tag-glow {
      background: rgba(255, 68, 68, 0.3);
      box-shadow: 0 0 15px rgba(255, 68, 68, 0.4);
    }
  }
  
  &--info {
    background: rgba(102, 221, 255, 0.2);
    color: #66ddff;
    border: 1px solid rgba(102, 221, 255, 0.4);
    
    .tag-glow {
      background: rgba(102, 221, 255, 0.3);
      box-shadow: 0 0 15px rgba(102, 221, 255, 0.4);
    }
  }
  
  // Glow effect
  &--glow {
    .tag-glow {
      opacity: 0.6;
      animation: tag-pulse 2s ease-in-out infinite;
    }
    
    &:hover {
      .tag-glow {
        opacity: 1;
      }
    }
  }
  
  // Hover effects
  &:hover {
    transform: translateY(-1px);
    
    .tag-glow {
      opacity: 0.8;
    }
  }
}

@keyframes tag-pulse {
  0%, 100% {
    opacity: 0.4;
  }
  50% {
    opacity: 0.8;
  }
}
</style> 