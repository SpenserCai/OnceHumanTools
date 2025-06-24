<template>
  <div class="hologram-radio-group">
    <div 
      v-for="option in options" 
      :key="option.value"
      class="radio-option"
      :class="{ 'radio-option--checked': modelValue === option.value }"
      @click="selectOption(option.value)"
    >
      <div class="radio-visual">
        <div class="radio-outer">
          <div class="radio-inner" v-if="modelValue === option.value"></div>
        </div>
        <div class="radio-glow"></div>
      </div>
      <span class="radio-label">{{ option.label }}</span>
    </div>
  </div>
</template>

<script setup>
const props = defineProps({
  modelValue: {
    required: true
  },
  options: {
    type: Array,
    required: true,
    default: () => []
  }
})

const emit = defineEmits(['update:modelValue'])

const selectOption = (value) => {
  emit('update:modelValue', value)
}
</script>

<style lang="scss" scoped>
@use '@/styles/variables' as *;

.hologram-radio-group {
  display: flex;
  flex-direction: column;
  gap: $spacing-md;
  
  .radio-option {
    display: flex;
    align-items: center;
    gap: $spacing-sm;
    cursor: pointer;
    padding: $spacing-sm;
    border-radius: $radius-sm;
    transition: all $transition-normal;
    
    &:hover {
      background: rgba(0, 212, 255, 0.05);
    }
    
    .radio-visual {
      position: relative;
      width: 20px;
      height: 20px;
      flex-shrink: 0;
      
      .radio-outer {
        width: 100%;
        height: 100%;
        border: 2px solid rgba(0, 212, 255, 0.4);
        border-radius: 50%;
        display: flex;
        align-items: center;
        justify-content: center;
        transition: all $transition-normal;
      }
      
      .radio-inner {
        width: 8px;
        height: 8px;
        background: $primary-color;
        border-radius: 50%;
        animation: radio-appear $transition-normal ease-out;
        filter: drop-shadow(0 0 4px currentColor);
      }
      
      .radio-glow {
        position: absolute;
        top: -2px;
        left: -2px;
        right: -2px;
        bottom: -2px;
        border-radius: 50%;
        background: radial-gradient(circle, rgba(0, 212, 255, 0.3) 0%, transparent 70%);
        opacity: 0;
        transition: opacity $transition-normal;
      }
    }
    
    .radio-label {
      color: $text-primary;
      font-size: 0.875rem;
      user-select: none;
    }
    
    // Checked state
    &--checked {
      .radio-outer {
        border-color: $primary-color;
        box-shadow: 0 0 10px rgba(0, 212, 255, 0.3);
      }
      
      .radio-glow {
        opacity: 0.8;
      }
      
      .radio-label {
        color: $primary-color;
      }
    }
    
    // Hover state
    &:hover:not(.radio-option--checked) {
      .radio-outer {
        border-color: rgba(0, 212, 255, 0.7);
      }
      
      .radio-glow {
        opacity: 0.4;
      }
    }
  }
}

@keyframes radio-appear {
  0% {
    opacity: 0;
    transform: scale(0.3);
  }
  50% {
    opacity: 1;
    transform: scale(1.2);
  }
  100% {
    opacity: 1;
    transform: scale(1);
  }
}
</style> 