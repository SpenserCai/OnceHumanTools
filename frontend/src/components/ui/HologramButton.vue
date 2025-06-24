<template>
  <button 
    :class="[
      'hologram-button',
      `hologram-button--${variant}`,
      `hologram-button--${size}`,
      {
        'hologram-button--disabled': disabled,
        'hologram-button--loading': loading
      }
    ]"
    :disabled="disabled || loading"
    @click="handleClick"
  >
    <div class="button-background"></div>
    <div class="button-content">
      <div v-if="loading" class="loading-spinner"></div>
      <slot v-else></slot>
    </div>
    <div class="button-glow"></div>
  </button>
</template>

<script setup>
const props = defineProps({
  variant: {
    type: String,
    default: 'primary',
    validator: (value) => ['primary', 'secondary', 'outline', 'ghost'].includes(value)
  },
  size: {
    type: String,
    default: 'medium',
    validator: (value) => ['small', 'medium', 'large'].includes(value)
  },
  disabled: {
    type: Boolean,
    default: false
  },
  loading: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['click'])

const handleClick = (event) => {
  if (!props.disabled && !props.loading) {
    emit('click', event)
  }
}
</script>

<style lang="scss" scoped>
@use '@/styles/variables' as *;

.hologram-button {
  position: relative;
  border: none;
  cursor: pointer;
  font-family: $font-tech;
  font-weight: 500;
  text-transform: uppercase;
  letter-spacing: 1px;
  transition: all $transition-normal;
  overflow: hidden;
  outline: none;
  
  &:disabled {
    cursor: not-allowed;
    opacity: 0.5;
  }
  
  .button-background {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    transition: all $transition-normal;
  }
  
  .button-content {
    position: relative;
    z-index: 2;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: $spacing-xs;
  }
  
  .button-glow {
    position: absolute;
    top: -2px;
    left: -2px;
    right: -2px;
    bottom: -2px;
    border-radius: inherit;
    opacity: 0;
    transition: opacity $transition-normal;
    pointer-events: none;
  }
  
  .loading-spinner {
    width: 16px;
    height: 16px;
    border: 2px solid rgba(255, 255, 255, 0.3);
    border-top: 2px solid currentColor;
    border-radius: 50%;
    animation: spin 1s linear infinite;
  }
  
  // Sizes
  &--small {
    padding: $spacing-xs $spacing-md;
    font-size: 0.75rem;
    border-radius: $radius-sm;
  }
  
  &--medium {
    padding: $spacing-sm $spacing-lg;
    font-size: 0.875rem;
    border-radius: $radius-md;
  }
  
  &--large {
    padding: $spacing-md $spacing-xl;
    font-size: 1rem;
    border-radius: $radius-lg;
  }
  
  // Variants
  &--primary {
    color: $text-primary;
    
    .button-background {
      background: $hologram-primary;
      border: 1px solid rgba(0, 212, 255, 0.6);
    }
    
    .button-glow {
      background: $hologram-primary;
      box-shadow: $hologram-glow;
    }
    
    &:hover:not(:disabled) {
      transform: translateY(-2px);
      
      .button-glow {
        opacity: 0.7;
      }
    }
    
    &:active:not(:disabled) {
      transform: translateY(0);
    }
  }
  
  &--secondary {
    color: $text-primary;
    
    .button-background {
      background: $hologram-secondary;
      border: 1px solid rgba(0, 153, 204, 0.6);
    }
    
    .button-glow {
      background: $hologram-secondary;
      box-shadow: 0 0 20px rgba(0, 153, 204, 0.4);
    }
    
    &:hover:not(:disabled) {
      transform: translateY(-2px);
      
      .button-glow {
        opacity: 0.7;
      }
    }
  }
  
  &--outline {
    color: $primary-color;
    
    .button-background {
      background: transparent;
      border: 2px solid $primary-color;
    }
    
    .button-glow {
      background: transparent;
      border: 2px solid $primary-color;
      box-shadow: 0 0 20px rgba(0, 212, 255, 0.3);
    }
    
    &:hover:not(:disabled) {
      color: $text-primary;
      
      .button-background {
        background: rgba(0, 212, 255, 0.1);
      }
      
      .button-glow {
        opacity: 1;
      }
    }
  }
  
  &--ghost {
    color: $primary-color;
    
    .button-background {
      background: transparent;
      border: 1px solid transparent;
    }
    
    .button-glow {
      background: rgba(0, 212, 255, 0.1);
      border: 1px solid $primary-color;
    }
    
    &:hover:not(:disabled) {
      .button-background {
        background: rgba(0, 212, 255, 0.05);
        border-color: rgba(0, 212, 255, 0.3);
      }
      
      .button-glow {
        opacity: 1;
      }
    }
  }
  
  &--disabled {
    .button-background {
      background: rgba(100, 100, 100, 0.2) !important;
      border-color: rgba(100, 100, 100, 0.3) !important;
    }
    
    color: $text-muted !important;
    
    &:hover {
      transform: none !important;
      
      .button-glow {
        opacity: 0 !important;
      }
    }
  }
}

@keyframes spin {
  from {
    transform: rotate(0deg);
  }
  to {
    transform: rotate(360deg);
  }
}
</style> 