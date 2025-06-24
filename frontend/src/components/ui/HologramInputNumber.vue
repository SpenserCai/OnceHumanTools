<template>
  <div class="hologram-input-number" :class="{ 'hologram-input-number--disabled': disabled }">
    <button 
      class="input-button decrease"
      :disabled="disabled || modelValue <= min"
      @click="decrease"
    >
      <svg viewBox="0 0 24 24" class="button-icon">
        <path d="M19 13H5v-2h14v2z" fill="currentColor"/>
      </svg>
    </button>
    
    <input
      type="number"
      class="input-field"
      :value="modelValue"
      :min="min"
      :max="max"
      :disabled="disabled"
      @input="handleInput"
      @blur="handleBlur"
      @focus="handleFocus"
    />
    
    <button 
      class="input-button increase"
      :disabled="disabled || modelValue >= max"
      @click="increase"
    >
      <svg viewBox="0 0 24 24" class="button-icon">
        <path d="M19 13h-6v6h-2v-6H5v-2h6V5h2v6h6v2z" fill="currentColor"/>
      </svg>
    </button>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const props = defineProps({
  modelValue: {
    type: Number,
    default: 0
  },
  min: {
    type: Number,
    default: 0
  },
  max: {
    type: Number,
    default: 100
  },
  step: {
    type: Number,
    default: 1
  },
  disabled: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:modelValue'])

const focused = ref(false)

const updateValue = (value) => {
  const numValue = Number(value)
  if (isNaN(numValue)) return
  
  const clampedValue = Math.max(props.min, Math.min(props.max, numValue))
  emit('update:modelValue', clampedValue)
}

const increase = () => {
  const newValue = props.modelValue + props.step
  if (newValue <= props.max) {
    updateValue(newValue)
  }
}

const decrease = () => {
  const newValue = props.modelValue - props.step
  if (newValue >= props.min) {
    updateValue(newValue)
  }
}

const handleInput = (event) => {
  updateValue(event.target.value)
}

const handleFocus = () => {
  focused.value = true
}

const handleBlur = () => {
  focused.value = false
}
</script>

<style lang="scss" scoped>
@use '@/styles/variables' as *;

.hologram-input-number {
  display: flex;
  align-items: center;
  border: 1px solid rgba(0, 212, 255, 0.3);
  border-radius: $radius-sm;
  background: rgba(0, 33, 66, 0.6);
  transition: all $transition-normal;
  overflow: hidden;
  
  &:hover {
    border-color: rgba(0, 212, 255, 0.6);
    box-shadow: 0 0 10px rgba(0, 212, 255, 0.2);
  }
  
  &:focus-within {
    border-color: $primary-color;
    box-shadow: 0 0 15px rgba(0, 212, 255, 0.4);
  }
  
  .input-button {
    width: 32px;
    height: 32px;
    background: transparent;
    border: none;
    color: rgba(0, 212, 255, 0.8);
    cursor: pointer;
    display: flex;
    align-items: center;
    justify-content: center;
    transition: all $transition-fast;
    
    &:hover:not(:disabled) {
      background: rgba(0, 212, 255, 0.1);
      color: $primary-color;
    }
    
    &:active:not(:disabled) {
      background: rgba(0, 212, 255, 0.2);
    }
    
    &:disabled {
      color: rgba(100, 100, 100, 0.5);
      cursor: not-allowed;
    }
    
    .button-icon {
      width: 16px;
      height: 16px;
    }
  }
  
  .input-field {
    flex: 1;
    height: 32px;
    padding: 0 $spacing-xs;
    background: transparent;
    border: none;
    color: $text-primary;
    font-family: $font-tech;
    text-align: center;
    outline: none;
    
    &::placeholder {
      color: $text-muted;
    }
    
    &::-webkit-outer-spin-button,
    &::-webkit-inner-spin-button {
      -webkit-appearance: none;
      margin: 0;
    }
    
    &[type=number] {
      -moz-appearance: textfield;
    }
  }
  
  &--disabled {
    opacity: 0.5;
    cursor: not-allowed;
    
    &:hover {
      border-color: rgba(0, 212, 255, 0.3) !important;
      box-shadow: none !important;
    }
  }
}
</style> 