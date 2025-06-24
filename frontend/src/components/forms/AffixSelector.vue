<template>
  <div class="affix-selector">
    <div class="selector-header" v-if="title">
      <h3 class="selector-title">{{ title }}</h3>
    </div>
    
    <div class="selector-content">
      <div class="affix-grid">
        <HologramCheckbox
          v-for="affix in affixList"
          :key="affix.id"
          :modelValue="selectedAffixes.includes(affix.id)"
          @update:modelValue="(checked) => toggleAffix(affix.id, checked)"
          class="affix-item"
        >
          <div class="affix-content">
            <span class="affix-name">{{ affix.name }}</span>
            <div class="affix-meta">
              <span class="affix-category" :class="`category-${affix.category}`">
                {{ getCategoryName(affix.category) }}
              </span>
            </div>
          </div>
        </HologramCheckbox>
      </div>
    </div>
    
    <div class="selector-actions" v-if="showQuickActions">
      <div class="quick-actions">
        <HologramButton 
          variant="ghost" 
          size="small" 
          @click="selectCategory('damage')"
        >
          伤害类
        </HologramButton>
        <HologramButton 
          variant="ghost" 
          size="small" 
          @click="selectCategory('defense')"
        >
          防御类
        </HologramButton>
        <HologramButton 
          variant="ghost" 
          size="small" 
          @click="selectCategory('utility')"
        >
          功能类
        </HologramButton>
        <HologramButton 
          variant="outline" 
          size="small" 
          @click="clearSelection"
        >
          清空
        </HologramButton>
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import HologramCheckbox from '@/components/ui/HologramCheckbox.vue'
import HologramButton from '@/components/ui/HologramButton.vue'

const props = defineProps({
  modelValue: {
    type: Array,
    default: () => []
  },
  affixList: {
    type: Array,
    default: () => []
  },
  title: {
    type: String,
    default: ''
  },
  showQuickActions: {
    type: Boolean,
    default: true
  }
})

const emit = defineEmits(['update:modelValue'])

const selectedAffixes = computed(() => props.modelValue)

const toggleAffix = (affixId, checked) => {
  const newSelection = [...selectedAffixes.value]
  
  if (checked) {
    if (!newSelection.includes(affixId)) {
      newSelection.push(affixId)
    }
  } else {
    const index = newSelection.indexOf(affixId)
    if (index > -1) {
      newSelection.splice(index, 1)
    }
  }
  
  emit('update:modelValue', newSelection)
}

const selectCategory = (category) => {
  const categoryAffixes = props.affixList
    .filter(affix => affix.category === category)
    .map(affix => affix.id)
  
  emit('update:modelValue', categoryAffixes)
}

const clearSelection = () => {
  emit('update:modelValue', [])
}

const getCategoryName = (category) => {
  const nameMap = {
    damage: '伤害',
    defense: '防御',
    utility: '功能'
  }
  return nameMap[category] || category
}
</script>

<style lang="scss" scoped>
@use '@/styles/variables' as *;

.affix-selector {
  .selector-header {
    margin-bottom: $spacing-lg;
    
    .selector-title {
      font-family: $font-tech;
      font-size: 1.2rem;
      color: $primary-color;
      margin: 0;
      text-transform: uppercase;
      letter-spacing: 1px;
    }
  }
  
  .selector-content {
    margin-bottom: $spacing-lg;
    
    .affix-grid {
      max-height: 300px;
      overflow-y: auto;
      padding: $spacing-md;
      background: rgba(0, 33, 66, 0.3);
      border: 1px solid $border-color;
      border-radius: $radius-md;
      position: relative;
      

      
      // 隐藏滚动条
      &::-webkit-scrollbar {
        width: 0px;
        display: none;
      }
      
      .affix-item {
        margin-bottom: $spacing-md;
        padding: $spacing-sm;
        border-radius: $radius-sm;
        transition: all $transition-normal;
        
        &:hover {
          background: rgba(0, 212, 255, 0.05);
        }
        
        .affix-content {
          display: flex;
          flex-direction: column;
          gap: $spacing-xs;
          
          .affix-name {
            font-weight: 500;
            color: $text-primary;
          }
          
          .affix-meta {
            .affix-category {
              display: inline-block;
              padding: 2px $spacing-xs;
              border-radius: $radius-sm;
              font-size: 0.75rem;
              font-weight: 500;
              text-transform: uppercase;
              letter-spacing: 0.5px;
              
              &.category-damage {
                background: rgba(255, 68, 68, 0.2);
                color: #ff6b6b;
                border: 1px solid rgba(255, 107, 107, 0.3);
              }
              
              &.category-defense {
                background: rgba(0, 255, 170, 0.2);
                color: #00ffaa;
                border: 1px solid rgba(0, 255, 170, 0.3);
              }
              
              &.category-utility {
                background: rgba(255, 170, 0, 0.2);
                color: #ffaa00;
                border: 1px solid rgba(255, 170, 0, 0.3);
              }
            }
          }
        }
      }
    }
  }
  
  .selector-actions {
    .quick-actions {
      display: flex;
      gap: $spacing-sm;
      flex-wrap: wrap;
      
      @media (max-width: 768px) {
        justify-content: center;
      }
    }
  }
}
</style> 