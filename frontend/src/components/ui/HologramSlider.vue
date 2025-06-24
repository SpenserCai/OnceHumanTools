<template>
  <div class="hologram-slider">
    <div class="slider-label" v-if="label">{{ label }}</div>
    <div class="slider-container">
      <div class="slider-track" ref="track" @click="handleTrackClick">
        <div 
          class="slider-fill" 
          :style="{ width: fillPercentage + '%' }"
        ></div>
        <div 
          class="slider-thumb" 
          :style="{ left: fillPercentage + '%' }"
          @mousedown="startDrag"
          @touchstart="startDrag"
        >
          <div class="thumb-core"></div>
          <div class="thumb-glow"></div>
        </div>
      </div>
      <div class="slider-marks" v-if="marks">
        <div 
          v-for="(mark, value) in marks" 
          :key="value"
          class="mark"
          :style="{ left: ((value - min) / (max - min)) * 100 + '%' }"
        >
          <div class="mark-dot"></div>
          <div class="mark-label">{{ mark }}</div>
        </div>
      </div>
    </div>
    <div class="slider-input" v-if="showInput">
      <input 
        type="number" 
        :value="modelValue" 
        @input="handleInputChange"
        :min="min"
        :max="max"
        class="hologram-input"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'

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
  label: {
    type: String,
    default: ''
  },
  marks: {
    type: Object,
    default: null
  },
  showInput: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['update:modelValue'])

const track = ref(null)
const isDragging = ref(false)

const fillPercentage = computed(() => {
  return ((props.modelValue - props.min) / (props.max - props.min)) * 100
})

const updateValue = (percentage) => {
  const value = props.min + (percentage / 100) * (props.max - props.min)
  const clampedValue = Math.max(props.min, Math.min(props.max, Math.round(value)))
  emit('update:modelValue', clampedValue)
}

const handleTrackClick = (event) => {
  if (isDragging.value) return
  
  const rect = track.value.getBoundingClientRect()
  const percentage = ((event.clientX - rect.left) / rect.width) * 100
  updateValue(percentage)
}

const startDrag = (event) => {
  event.preventDefault()
  isDragging.value = true
  
  const handleDrag = (e) => {
    if (!isDragging.value) return
    
    const rect = track.value.getBoundingClientRect()
    const clientX = e.clientX || (e.touches && e.touches[0].clientX)
    const percentage = Math.max(0, Math.min(100, ((clientX - rect.left) / rect.width) * 100))
    updateValue(percentage)
  }
  
  const stopDrag = () => {
    isDragging.value = false
    document.removeEventListener('mousemove', handleDrag)
    document.removeEventListener('mouseup', stopDrag)
    document.removeEventListener('touchmove', handleDrag)
    document.removeEventListener('touchend', stopDrag)
  }
  
  document.addEventListener('mousemove', handleDrag)
  document.addEventListener('mouseup', stopDrag)
  document.addEventListener('touchmove', handleDrag)
  document.addEventListener('touchend', stopDrag)
}

const handleInputChange = (event) => {
  const value = parseInt(event.target.value)
  if (!isNaN(value)) {
    emit('update:modelValue', Math.max(props.min, Math.min(props.max, value)))
  }
}
</script>

<style lang="scss" scoped>
@use '@/styles/variables' as *;

.hologram-slider {
  .slider-label {
    color: $text-secondary;
    font-size: 0.9rem;
    margin-bottom: $spacing-sm;
    font-weight: 500;
  }
  
  .slider-container {
    position: relative;
    margin-bottom: $spacing-md;
  }
  
  .slider-track {
    position: relative;
    height: 6px;
    background: rgba(0, 212, 255, 0.1);
    border-radius: 3px;
    border: 1px solid rgba(0, 212, 255, 0.2);
    cursor: pointer;
    overflow: visible;
    
    &::before {
      content: '';
      position: absolute;
      top: -1px;
      left: -1px;
      right: -1px;
      bottom: -1px;
      background: $hologram-border;
      border-radius: 3px;
      opacity: 0;
      transition: opacity $transition-normal;
    }
    
    &:hover::before {
      opacity: 0.5;
    }
  }
  
  .slider-fill {
    height: 100%;
    background: $hologram-primary;
    border-radius: 3px;
    position: relative;
    transition: width $transition-fast;
    
    &::after {
      content: '';
      position: absolute;
      top: 0;
      left: 0;
      right: 0;
      bottom: 0;
      background: $hologram-primary;
      border-radius: 3px;
      animation: hologram-pulse 2s ease-in-out infinite;
    }
  }
  
  .slider-thumb {
    position: absolute;
    top: 50%;
    transform: translate(-50%, -50%);
    width: 20px;
    height: 20px;
    cursor: pointer;
    z-index: 10;
    
    .thumb-core {
      width: 12px;
      height: 12px;
      background: $primary-color;
      border-radius: 50%;
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      border: 2px solid rgba(0, 212, 255, 0.8);
      transition: all $transition-fast;
    }
    
    .thumb-glow {
      position: absolute;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      width: 20px;
      height: 20px;
      border-radius: 50%;
      background: radial-gradient(circle, rgba(0, 212, 255, 0.4) 0%, transparent 70%);
      opacity: 0;
      transition: opacity $transition-fast;
    }
    
    &:hover {
      .thumb-core {
        transform: translate(-50%, -50%) scale(1.2);
        box-shadow: $hologram-glow;
      }
      
      .thumb-glow {
        opacity: 1;
      }
    }
  }
  
  .slider-marks {
    position: absolute;
    top: 20px;
    left: 0;
    right: 0;
    height: 20px;
    
    .mark {
      position: absolute;
      transform: translateX(-50%);
      
      .mark-dot {
        width: 4px;
        height: 4px;
        background: rgba(0, 212, 255, 0.6);
        border-radius: 50%;
        margin: 0 auto 4px;
      }
      
      .mark-label {
        font-size: 0.75rem;
        color: $text-muted;
        text-align: center;
        min-width: 20px;
      }
    }
  }
  
  .slider-input {
    .hologram-input {
      width: 80px;
      padding: $spacing-xs $spacing-sm;
      background: rgba(0, 33, 66, 0.6);
      border: 1px solid rgba(0, 212, 255, 0.3);
      border-radius: $radius-sm;
      color: $text-primary;
      font-family: $font-tech;
      text-align: center;
      outline: none;
      transition: all $transition-normal;
      
      &:focus {
        border-color: $primary-color;
        box-shadow: 0 0 10px rgba(0, 212, 255, 0.3);
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
  }
}

@keyframes hologram-pulse {
  0%, 100% {
    opacity: 0.8;
  }
  50% {
    opacity: 1;
  }
}
</style> 