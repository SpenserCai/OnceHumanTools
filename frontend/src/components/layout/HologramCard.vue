<template>
  <div 
    :class="[
      'hologram-card',
      `hologram-card--${variant}`,
      {
        'hologram-card--interactive': interactive,
        'hologram-card--glow': glow
      }
    ]"
    @click="handleClick"
  >
    <div class="card-background"></div>
    <div class="card-border"></div>
    <div class="card-glow"></div>
    
    <div class="card-content">
      <div class="card-header" v-if="title || $slots.header">
        <slot name="header">
          <h3 class="card-title" v-if="title">{{ title }}</h3>
        </slot>
      </div>
      
      <div class="card-body">
        <slot></slot>
      </div>
      
      <div class="card-footer" v-if="$slots.footer">
        <slot name="footer"></slot>
      </div>
    </div>
    
    <div class="card-corner-effects">
      <div class="corner corner-tl"></div>
      <div class="corner corner-tr"></div>
      <div class="corner corner-bl"></div>
      <div class="corner corner-br"></div>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  title: {
    type: String,
    default: ''
  },
  variant: {
    type: String,
    default: 'default',
    validator: (value) => ['default', 'primary', 'secondary', 'accent'].includes(value)
  },
  interactive: {
    type: Boolean,
    default: false
  },
  glow: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['click'])

const handleClick = (event) => {
  if (props.interactive) {
    emit('click', event)
  }
}
</script>

<style lang="scss" scoped>
@use '@/styles/variables' as *;

.hologram-card {
  position: relative;
  border-radius: $radius-lg;
  overflow: hidden;
  transition: all $transition-normal;
  
  .card-background {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    background: $bg-card;
    backdrop-filter: blur(10px);
    transition: all $transition-normal;
  }
  
  .card-border {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    border: 1px solid $border-color;
    border-radius: $radius-lg;
    transition: all $transition-normal;
  }
  
  .card-glow {
    position: absolute;
    top: -2px;
    left: -2px;
    right: -2px;
    bottom: -2px;
    border-radius: $radius-lg;
    opacity: 0;
    transition: opacity $transition-normal;
    pointer-events: none;
  }
  
  .card-content {
    position: relative;
    z-index: 2;
    padding: $spacing-lg;
    
    .card-header {
      margin-bottom: $spacing-lg;
      
      .card-title {
        font-family: $font-tech;
        font-size: 1.5rem;
        color: $primary-color;
        margin: 0;
        text-transform: uppercase;
        letter-spacing: 1px;
        text-shadow: 0 0 10px rgba(0, 212, 255, 0.3);
      }
    }
    
    .card-body {
      color: $text-secondary;
    }
    
    .card-footer {
      margin-top: $spacing-lg;
      padding-top: $spacing-md;
      border-top: 1px solid rgba(0, 212, 255, 0.2);
    }
  }
  
  .card-corner-effects {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    pointer-events: none;
    
    .corner {
      position: absolute;
      width: 20px;
      height: 20px;
      border: 2px solid transparent;
      transition: all $transition-normal;
      
      &.corner-tl {
        top: 0;
        left: 0;
        border-top-color: $primary-color;
        border-left-color: $primary-color;
        border-top-left-radius: $radius-lg;
      }
      
      &.corner-tr {
        top: 0;
        right: 0;
        border-top-color: $primary-color;
        border-right-color: $primary-color;
        border-top-right-radius: $radius-lg;
      }
      
      &.corner-bl {
        bottom: 0;
        left: 0;
        border-bottom-color: $primary-color;
        border-left-color: $primary-color;
        border-bottom-left-radius: $radius-lg;
      }
      
      &.corner-br {
        bottom: 0;
        right: 0;
        border-bottom-color: $primary-color;
        border-right-color: $primary-color;
        border-bottom-right-radius: $radius-lg;
      }
    }
  }
  
  // Variants
  &--default {
    .card-glow {
      background: $hologram-primary;
      box-shadow: 0 0 30px rgba(0, 212, 255, 0.2);
    }
  }
  
  &--primary {
    .card-border {
      border-color: $primary-color;
    }
    
    .card-glow {
      background: $hologram-primary;
      box-shadow: $hologram-glow;
    }
  }
  
  &--secondary {
    .card-border {
      border-color: $secondary-color;
    }
    
    .card-glow {
      background: $hologram-secondary;
      box-shadow: 0 0 30px rgba(0, 153, 204, 0.4);
    }
    
    .card-title {
      color: $secondary-color;
    }
    
    .corner {
      border-color: $secondary-color !important;
    }
  }
  
  &--accent {
    .card-border {
      border-color: $accent-color;
    }
    
    .card-glow {
      background: linear-gradient(45deg, $accent-color, rgba(102, 221, 255, 0.8));
      box-shadow: 0 0 30px rgba(102, 221, 255, 0.4);
    }
    
    .card-title {
      color: $accent-color;
    }
    
    .corner {
      border-color: $accent-color !important;
    }
  }
  
  // States
  &--interactive {
    cursor: pointer;
    
    &:hover {
      transform: translateY(-4px);
      
      .card-border {
        border-color: $primary-color;
        box-shadow: 0 8px 25px rgba(0, 212, 255, 0.2);
      }
      
      .card-glow {
        opacity: 0.8;
      }
      
      .corner {
        width: 30px;
        height: 30px;
      }
    }
    
    &:active {
      transform: translateY(-2px);
    }
  }
  
  &--glow {
    .card-glow {
      opacity: 0.6;
      animation: hologram-pulse 3s ease-in-out infinite;
    }
  }
  
  // Hover effects
  &:hover {
    .card-background {
      background: rgba(0, 33, 66, 0.9);
    }
  }
}

@keyframes hologram-pulse {
  0%, 100% {
    opacity: 0.4;
  }
  50% {
    opacity: 0.8;
  }
}
</style> 