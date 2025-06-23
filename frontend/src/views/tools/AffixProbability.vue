<template>
  <div class="affix-probability-page">
    <div class="page-header">
      <h1 class="page-title glow-text">模组词条概率计算器</h1>
      <p class="page-desc">计算特定词条组合出现的概率</p>
    </div>
    
    <div class="tool-container">
      <!-- 输入区域 -->
      <div class="input-section sci-fi-card">
        <h2 class="section-title">参数设置</h2>
        
        <!-- 词条数量选择 -->
        <div class="form-group">
          <label class="form-label">词条数量</label>
          <el-slider 
            v-model="slotCount" 
            :min="1" 
            :max="10"
            :marks="slotMarks"
            show-input
          />
        </div>
        
        <!-- 词条选择 -->
        <div class="form-group">
          <label class="form-label">目标词条范围</label>
          <div class="affix-selector">
            <el-checkbox-group v-model="selectedAffixes">
              <div 
                v-for="affix in affixList" 
                :key="affix.id"
                class="affix-item"
              >
                <el-checkbox 
                  :label="affix.id"
                  :disabled="selectedAffixes.length >= slotCount && !selectedAffixes.includes(affix.id)"
                >
                  <span class="affix-name">{{ affix.name }}</span>
                  <el-tag 
                    :type="getTagType(affix.category)" 
                    size="small"
                    class="affix-tag"
                  >
                    {{ getCategoryName(affix.category) }}
                  </el-tag>
                </el-checkbox>
              </div>
            </el-checkbox-group>
          </div>
        </div>
        
        <!-- 快速选择 -->
        <div class="form-group">
          <label class="form-label">快速选择</label>
          <div class="quick-select">
            <el-button @click="selectCategory('damage')">伤害类</el-button>
            <el-button @click="selectCategory('defense')">防御类</el-button>
            <el-button @click="selectCategory('utility')">功能类</el-button>
            <el-button @click="clearSelection">清空</el-button>
          </div>
        </div>
        
        <!-- 显示选项 -->
        <div class="form-group">
          <el-checkbox v-model="showCombinations">
            显示具体组合
          </el-checkbox>
        </div>
        
        <!-- 计算按钮 -->
        <div class="form-actions">
          <button 
            class="sci-fi-btn"
            :disabled="selectedAffixes.length === 0"
            @click="calculate"
          >
            开始计算
          </button>
        </div>
      </div>
      
      <!-- 结果区域 -->
      <div v-if="result" class="result-section sci-fi-card fade-in">
        <h2 class="section-title">计算结果</h2>
        
        <!-- 概率显示 -->
        <div class="probability-display">
          <div class="probability-value glow-text">
            {{ (result.probabilityPercent || 0).toFixed(4) }}%
          </div>
          <div class="probability-label">出现概率</div>
        </div>
        
        <!-- 详细数据 -->
        <div class="result-stats">
          <div class="stat-item">
            <span class="stat-label">满足条件的组合数</span>
            <span class="stat-value">{{ result.validCombinations }}</span>
          </div>
          <div class="stat-item">
            <span class="stat-label">总组合数</span>
            <span class="stat-value">{{ result.totalCombinations }}</span>
          </div>
          <div class="stat-item">
            <span class="stat-label">精确概率</span>
            <span class="stat-value">{{ result.probability?.toFixed(6) || 0 }}</span>
          </div>
        </div>
        
        <!-- 组合列表 -->
        <div v-if="showCombinations && result.combinations" class="combinations-section">
          <h3 class="subsection-title">所有可能的组合</h3>
          <div class="combinations-grid">
            <div 
              v-for="(combo, index) in result.combinations" 
              :key="index"
              class="combination-item"
            >
              <span class="combo-index">#{{ index + 1 }}</span>
              <div class="combo-affixes">
                <el-tag 
                  v-for="affixId in combo" 
                  :key="affixId"
                  :type="getAffixTagType(affixId)"
                  effect="plain"
                >
                  {{ getAffixName(affixId) }}
                </el-tag>
              </div>
            </div>
          </div>
        </div>
        
        <!-- 图表展示 -->
        <div class="chart-section">
          <canvas ref="chartCanvas"></canvas>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { Chart, registerables } from 'chart.js'
import api from '@/api'

Chart.register(...registerables)

// 数据状态
const slotCount = ref(3)
const selectedAffixes = ref([])
const showCombinations = ref(false)
const affixList = ref([])
const result = ref(null)
const chartCanvas = ref(null)
let chartInstance = null

// 词条槽位标记
const slotMarks = computed(() => {
  const marks = {}
  for (let i = 1; i <= 10; i++) {
    marks[i] = i.toString()
  }
  return marks
})

// 获取词条列表
const fetchAffixList = async () => {
  try {
    const res = await api.mod.getAffixList()
    affixList.value = res.affixes || []
  } catch (error) {
    ElMessage.error('获取词条列表失败')
  }
}

// 计算概率
const calculate = async () => {
  if (selectedAffixes.value.length === 0) {
    ElMessage.warning('请选择至少一个目标词条')
    return
  }
  
  try {
    const res = await api.mod.calculateAffixProbability({
      slotCount: slotCount.value,
      targetAffixIds: selectedAffixes.value,
      showCombinations: showCombinations.value
    })
    
    result.value = res
    updateChart()
  } catch (error) {
    ElMessage.error('计算失败，请重试')
  }
}

// 更新图表
const updateChart = () => {
  if (!chartCanvas.value || !result.value) return
  
  // 销毁旧图表
  if (chartInstance) {
    chartInstance.destroy()
  }
  
  // 创建新图表
  const ctx = chartCanvas.value.getContext('2d')
  chartInstance = new Chart(ctx, {
    type: 'doughnut',
    data: {
      labels: ['目标组合', '其他组合'],
      datasets: [{
        data: [
          result.value.validCombinations,
          result.value.totalCombinations - result.value.validCombinations
        ],
        backgroundColor: ['#00ff88', '#333333'],
        borderColor: ['#00ff88', '#333333'],
        borderWidth: 2
      }]
    },
    options: {
      responsive: true,
      maintainAspectRatio: false,
      plugins: {
        legend: {
          position: 'bottom',
          labels: {
            color: '#fff',
            font: {
              family: 'Orbitron'
            }
          }
        },
        tooltip: {
          callbacks: {
            label: (context) => {
              const percentage = (context.parsed / result.value.totalCombinations * 100).toFixed(2)
              return `${context.label}: ${context.parsed} (${percentage}%)`
            }
          }
        }
      }
    }
  })
}

// 分类选择
const selectCategory = (category) => {
  const categoryAffixes = affixList.value
    .filter(affix => affix.category === category)
    .map(affix => affix.id)
  
  selectedAffixes.value = categoryAffixes.slice(0, slotCount.value)
}

// 清空选择
const clearSelection = () => {
  selectedAffixes.value = []
  result.value = null
}

// 获取词条名称
const getAffixName = (id) => {
  const affix = affixList.value.find(a => a.id === id)
  return affix?.name || ''
}

// 获取标签类型
const getTagType = (category) => {
  const typeMap = {
    damage: 'danger',
    defense: 'success',
    utility: 'warning'
  }
  return typeMap[category] || 'info'
}

// 获取分类名称
const getCategoryName = (category) => {
  const nameMap = {
    damage: '伤害',
    defense: '防御',
    utility: '功能'
  }
  return nameMap[category] || category
}

// 获取词条标签类型
const getAffixTagType = (id) => {
  const affix = affixList.value.find(a => a.id === id)
  return affix ? getTagType(affix.category) : 'info'
}

// 监听词条数量变化
watch(slotCount, () => {
  // 如果选中的词条超过了槽位数，截断
  if (selectedAffixes.value.length > slotCount.value) {
    selectedAffixes.value = selectedAffixes.value.slice(0, slotCount.value)
  }
})

onMounted(() => {
  fetchAffixList()
})
</script>

<style lang="scss" scoped>
@use '@/styles/variables' as *;

.affix-probability-page {
  min-height: 100vh;
  padding: 80px $spacing-lg $spacing-xxl;
  
  .page-header {
    text-align: center;
    margin-bottom: $spacing-xxl;
    
    .page-title {
      font-family: $font-tech;
      font-size: 2.5rem;
      margin-bottom: $spacing-md;
      text-transform: uppercase;
      letter-spacing: 2px;
    }
    
    .page-desc {
      color: $text-secondary;
      font-size: 1.1rem;
    }
  }
  
  .tool-container {
    max-width: 1200px;
    margin: 0 auto;
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: $spacing-xl;
    
    @media (max-width: 968px) {
      grid-template-columns: 1fr;
    }
  }
  
  .section-title {
    font-family: $font-tech;
    font-size: 1.5rem;
    margin-bottom: $spacing-lg;
    color: $primary-color;
  }
  
  .form-group {
    margin-bottom: $spacing-lg;
    
    .form-label {
      display: block;
      margin-bottom: $spacing-sm;
      color: $text-secondary;
      font-weight: 500;
    }
  }
  
  .affix-selector {
    max-height: 300px;
    overflow-y: auto;
    padding: $spacing-md;
    background: rgba(0, 0, 0, 0.3);
    border: 1px solid $border-color;
    border-radius: $radius-md;
    
    .affix-item {
      margin-bottom: $spacing-sm;
      
      .affix-name {
        margin-right: $spacing-sm;
      }
      
      .affix-tag {
        font-size: 0.75rem;
      }
    }
  }
  
  .quick-select {
    display: flex;
    gap: $spacing-sm;
    flex-wrap: wrap;
  }
  
  .form-actions {
    margin-top: $spacing-xl;
    text-align: center;
  }
  
  .result-section {
    .probability-display {
      text-align: center;
      margin-bottom: $spacing-xl;
      
      .probability-value {
        font-family: $font-tech;
        font-size: 4rem;
        font-weight: 900;
        line-height: 1;
      }
      
      .probability-label {
        color: $text-secondary;
        margin-top: $spacing-sm;
      }
    }
    
    .result-stats {
      display: grid;
      grid-template-columns: 1fr;
      gap: $spacing-md;
      margin-bottom: $spacing-xl;
      
      .stat-item {
        display: flex;
        justify-content: space-between;
        padding: $spacing-md;
        background: rgba(0, 0, 0, 0.3);
        border: 1px solid $border-color;
        border-radius: $radius-sm;
        
        .stat-label {
          color: $text-secondary;
        }
        
        .stat-value {
          font-family: $font-tech;
          color: $primary-color;
          font-weight: 500;
        }
      }
    }
    
    .combinations-section {
      margin-top: $spacing-xl;
      
      .subsection-title {
        font-size: 1.2rem;
        margin-bottom: $spacing-md;
        color: $text-primary;
      }
      
      .combinations-grid {
        max-height: 400px;
        overflow-y: auto;
        padding: $spacing-md;
        background: rgba(0, 0, 0, 0.3);
        border: 1px solid $border-color;
        border-radius: $radius-md;
        
        .combination-item {
          display: flex;
          align-items: center;
          gap: $spacing-md;
          padding: $spacing-sm;
          margin-bottom: $spacing-sm;
          
          &:hover {
            background: $bg-hover;
          }
          
          .combo-index {
            color: $text-muted;
            font-family: $font-tech;
            min-width: 40px;
          }
          
          .combo-affixes {
            display: flex;
            flex-wrap: wrap;
            gap: $spacing-xs;
          }
        }
      }
    }
    
    .chart-section {
      margin-top: $spacing-xl;
      height: 300px;
      
      canvas {
        width: 100% !important;
        height: 100% !important;
      }
    }
  }
}

// Element Plus 样式覆盖
:deep(.el-slider) {
  .el-slider__runway {
    background: rgba(0, 255, 136, 0.2);
  }
  
  .el-slider__bar {
    background: $primary-color;
  }
  
  .el-slider__button {
    border-color: $primary-color;
  }
}

:deep(.el-checkbox) {
  .el-checkbox__label {
    color: $text-primary;
  }
  
  .el-checkbox__input.is-checked {
    .el-checkbox__inner {
      background: $primary-color;
      border-color: $primary-color;
    }
  }
}
</style>