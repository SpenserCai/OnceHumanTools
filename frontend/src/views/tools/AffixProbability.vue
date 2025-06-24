<template>
  <div class="affix-probability-page">
    <div class="page-header">
      <h1 class="page-title glow-text">模组词条概率计算器</h1>
      <p class="page-desc">计算特定词条组合出现的概率</p>
    </div>
    
    <div class="tool-container">
      <!-- 输入区域 -->
      <HologramCard class="input-section" title="参数设置" variant="primary">
        
        <!-- 词条数量选择 -->
        <div class="form-group">
          <HologramSlider
            v-model="slotCount"
            :min="1"
            :max="10"
            :marks="slotMarks"
            :showInput="true"
            label="词条数量"
          />
        </div>
        
        <!-- 词条选择 -->
        <div class="form-group">
          <AffixSelector
            v-model="selectedAffixes"
            :affixList="affixList"
            title="目标词条范围"
            :showQuickActions="true"
          />
        </div>
        
        <!-- 显示选项 -->
        <div class="form-group">
          <HologramCheckbox v-model="showCombinations">
            显示具体组合
          </HologramCheckbox>
        </div>
        
        <!-- 计算按钮 -->
        <div class="form-actions">
          <HologramButton
            variant="primary"
            size="large"
            :disabled="selectedAffixes.length === 0"
            @click="calculate"
          >
            开始计算
          </HologramButton>
        </div>
      </HologramCard>
      
      <!-- 结果区域 -->
      <HologramCard v-if="result" class="result-section fade-in" title="计算结果" variant="secondary" glow>
        
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
      </HologramCard>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { ElMessage } from 'element-plus'
import { Chart, registerables } from 'chart.js'
import api from '@/api'
import HologramSlider from '@/components/ui/HologramSlider.vue'
import HologramButton from '@/components/ui/HologramButton.vue'
import HologramCheckbox from '@/components/ui/HologramCheckbox.vue'
import AffixSelector from '@/components/forms/AffixSelector.vue'
import HologramCard from '@/components/layout/HologramCard.vue'

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
        backgroundColor: ['#00d4ff', '#333333'],
        borderColor: ['#00d4ff', '#333333'],
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

// 这些功能现在由AffixSelector组件处理

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
  // 词条数量变化时，不再限制目标词条范围的数量
  // 用户可以选择任意数量的目标词条作为候选范围
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
  
  // affix-selector 样式现在在专门的组件中处理
  
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

// 不再需要Element Plus样式覆盖，因为我们使用了自定义组件
</style>