<template>
  <label class="hologram-checkbox" :class="{ 'hologram-checkbox--disabled': disabled }">
    <input
      type="checkbox"
      :checked="modelValue"
      :disabled="disabled"
      @change="handleChange"
      class="checkbox-input"
    />
    <div class="checkbox-visual">
      <div class="checkbox-box">
        <div class="checkbox-background"></div>
        <div class="checkbox-border"></div>
        <div class="checkbox-checkmark" v-if="modelValue">
          <svg viewBox="0 0 12 12" class="checkmark-icon">
            <polyline points="1,6 4,9 11,2" stroke="currentColor" stroke-width="2" fill="none"/>
          </svg>
        </div>
        <div class="checkbox-glow"></div>
      </div>
    </div>
    <div class="checkbox-content" v-if="$slots.default">
      <slot></slot>
    </div>
  </label>
</template>

<script setup>
const props = defineProps({
  modelValue: {
    type: Boolean,
    default: false
  },
  disabled: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:modelValue'])

const handleChange = (event) => {
  if (!props.disabled) {
    emit('update:modelValue', event.target.checked)
  }
}
</script>

<style lang="scss" scoped>
@use '@/styles/variables' as *;

.hologram-checkbox {
  display: flex;
  align-items: center;
  gap: $spacing-sm;
  cursor: pointer;
  user-select: none;
  transition: all $transition-normal;
  
  &--disabled {
    cursor: not-allowed;
    opacity: 0.5;
  }
  
  .checkbox-input {
    position: absolute;
    opacity: 0;
    pointer-events: none;
  }
  
  .checkbox-visual {
    flex-shrink: 0;
  }
  
  .checkbox-box {
    position: relative;
    width: 18px;
    height: 18px;
    
    .checkbox-background {
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      background: rgba(0, 33, 66, 0.6);
      border-radius: $radius-sm;
      transition: all $transition-normal;
    }
    
    .checkbox-border {
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      border: 2px solid rgba(0, 212, 255, 0.4);
      border-radius: $radius-sm;
      transition: all $transition-normal;
    }
    
    .checkbox-checkmark {
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      width: 12px;
      height: 12px;
      color: $primary-color;
      animation: checkmark-appear $transition-normal ease-out;
      
      .checkmark-icon {
        width: 100%;
        height: 100%;
        filter: drop-shadow(0 0 4px currentColor);
      }
    }
    
    .checkbox-glow {
      position: absolute;
      top: -2px;
      left: -2px;
      right: -2px;
      bottom: -2px;
      background: radial-gradient(circle, rgba(0, 212, 255, 0.3) 0%, transparent 70%);
      border-radius: $radius-sm;
      opacity: 0;
      transition: opacity $transition-normal;
    }
  }
  
  .checkbox-content {
    color: $text-primary;
    font-size: 0.875rem;
    line-height: 1.4;
  }
  
  // States
  &:hover:not(.hologram-checkbox--disabled) {
    .checkbox-border {
      border-color: rgba(0, 212, 255, 0.7);
      box-shadow: 0 0 8px rgba(0, 212, 255, 0.2);
    }
    
    .checkbox-glow {
      opacity: 0.6;
    }
  }
  
  // Checked state
  .checkbox-input:checked + .checkbox-visual .checkbox-box {
    .checkbox-background {
      background: $hologram-primary;
    }
    
    .checkbox-border {
      border-color: $primary-color;
      box-shadow: 0 0 12px rgba(0, 212, 255, 0.4);
    }
    
    .checkbox-glow {
      opacity: 0.8;
    }
  }
  
  // Focus state
  .checkbox-input:focus + .checkbox-visual .checkbox-box {
    .checkbox-border {
      border-color: $primary-color;
      box-shadow: 0 0 0 3px rgba(0, 212, 255, 0.2);
    }
  }
  
  // Disabled state
  &.hologram-checkbox--disabled {
    .checkbox-background {
      background: rgba(100, 100, 100, 0.2) !important;
    }
    
    .checkbox-border {
      border-color: rgba(100, 100, 100, 0.3) !important;
    }
    
    .checkbox-content {
      color: $text-muted !important;
    }
    
    .checkbox-checkmark {
      color: $text-muted !important;
    }
    
    &:hover {
      .checkbox-glow {
        opacity: 0 !important;
      }
    }
  }
}

@keyframes checkmark-appear {
  0% {
    opacity: 0;
    transform: translate(-50%, -50%) scale(0.3);
  }
  50% {
    opacity: 1;
    transform: translate(-50%, -50%) scale(1.2);
  }
  100% {
    opacity: 1;
    transform: translate(-50%, -50%) scale(1);
  }
}
</style> 