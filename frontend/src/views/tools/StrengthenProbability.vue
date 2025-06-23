<template>
  <div class="strengthen-probability-page">
    <div class="page-header">
      <h1 class="page-title glow-text">模组强化概率计算器</h1>
      <p class="page-desc">计算模组词条强化到目标等级的成功率</p>
    </div>
    
    <div class="tool-container">
      <!-- 输入区域 -->
      <div class="input-section sci-fi-card">
        <h2 class="section-title">参数设置</h2>
        
        <!-- 初始等级 -->
        <div class="form-group">
          <label class="form-label">初始等级</label>
          <div class="level-inputs">
            <el-input-number
              v-for="(level, index) in initialLevels"
              :key="`initial-${index}`"
              v-model="initialLevels[index]"
              :min="1"
              :max="5"
              :precision="0"
              controls-position="right"
            />
          </div>
        </div>
        
        <!-- 目标等级 -->
        <div class="form-group">
          <label class="form-label">目标等级</label>
          <div class="level-inputs">
            <el-input-number
              v-for="(level, index) in targetLevels"
              :key="`target-${index}`"
              v-model="targetLevels[index]"
              :min="initialLevels[index]"
              :max="5"
              :precision="0"
              controls-position="right"
            />
          </div>
        </div>
        
        <!-- 判断模式 -->
        <div class="form-group">
          <label class="form-label">判断模式</label>
          <el-radio-group v-model="orderIndependent">
            <el-radio :label="true">顺序无关（推荐）</el-radio>
            <el-radio :label="false">位置对应</el-radio>
          </el-radio-group>
        </div>
        
        <!-- 显示选项 -->
        <div class="form-group">
          <el-checkbox v-model="showPaths">
            显示强化路径（数量较多时可能影响性能）
          </el-checkbox>
        </div>
        
        <!-- 预设方案 -->
        <div class="form-group">
          <label class="form-label">快速预设</label>
          <div class="preset-buttons">
            <el-button @click="applyPreset('allOne')">全1级</el-button>
            <el-button @click="applyPreset('balanced')">平衡型</el-button>
            <el-button @click="applyPreset('focused')">集中型</el-button>
          </div>
        </div>
        
        <!-- 计算按钮 -->
        <div class="form-actions">
          <button 
            class="sci-fi-btn"
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
          <div class="probability-value" :class="getProbabilityClass()">
            {{ (result.probabilityPercent || 0).toFixed(4) }}%
          </div>
          <div class="probability-label">成功概率</div>
        </div>
        
        <!-- 提示信息 -->
        <div class="probability-hint">
          <el-tag 
            :type="getProbabilityTagType()"
            effect="plain"
            size="large"
          >
            {{ getProbabilityHint() }}
          </el-tag>
        </div>
        
        <!-- 详细数据 -->
        <div class="result-stats">
          <div class="stat-item">
            <span class="stat-label">成功路径数</span>
            <span class="stat-value">{{ result.successfulOutcomes }}</span>
          </div>
          <div class="stat-item">
            <span class="stat-label">总路径数</span>
            <span class="stat-value">{{ result.totalOutcomes }}</span>
          </div>
          <div class="stat-item">
            <span class="stat-label">判断模式</span>
            <span class="stat-value">{{ orderIndependent ? '顺序无关' : '位置对应' }}</span>
          </div>
        </div>
        
        <!-- 路径展示 -->
        <div v-if="showPaths && result.paths && result.paths.length > 0" class="paths-section">
          <h3 class="subsection-title">强化路径示例（前10条）</h3>
          <div class="paths-list">
            <div 
              v-for="(path, index) in result.paths.slice(0, 10)" 
              :key="index"
              class="path-item"
              :class="{ success: path.success }"
            >
              <span class="path-index">#{{ index + 1 }}</span>
              <span class="path-result">{{ path.success ? '成功' : '失败' }}</span>
              <span class="path-final">最终: [{{ path.finalLevels.join(', ') }}]</span>
            </div>
          </div>
        </div>
        
        <!-- 可视化 -->
        <div class="visualization-section">
          <h3 class="subsection-title">概率可视化</h3>
          <div class="bar-chart">
            <div class="bar success-bar" :style="{ width: `${Math.min(result.probabilityPercent, 100)}%` }">
              <span v-if="result.probabilityPercent > 10">成功</span>
            </div>
            <div class="bar fail-bar" :style="{ width: `${Math.min(100 - result.probabilityPercent, 100)}%` }">
              <span v-if="100 - result.probabilityPercent > 10">失败</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import api from '@/api'

// 数据状态
const initialLevels = ref([1, 1, 1, 1])
const targetLevels = ref([2, 2, 2, 2])
const orderIndependent = ref(true)
const showPaths = ref(false)
const result = ref(null)

// 计算概率
const calculate = async () => {
  // 验证输入
  for (let i = 0; i < 4; i++) {
    if (targetLevels.value[i] < initialLevels.value[i]) {
      ElMessage.warning(`目标等级不能低于初始等级（词条${i + 1}）`)
      return
    }
  }
  
  try {
    const res = await api.mod.calculateStrengthenProbability({
      initialLevels: initialLevels.value,
      targetLevels: targetLevels.value,
      orderIndependent: orderIndependent.value,
      showPaths: showPaths.value
    })
    
    result.value = res
  } catch (error) {
    ElMessage.error('计算失败，请重试')
  }
}

// 应用预设
const applyPreset = (type) => {
  switch (type) {
    case 'allOne':
      initialLevels.value = [1, 1, 1, 1]
      targetLevels.value = [2, 2, 2, 2]
      break
    case 'balanced':
      initialLevels.value = [2, 2, 2, 2]
      targetLevels.value = [3, 3, 3, 3]
      break
    case 'focused':
      initialLevels.value = [1, 1, 1, 1]
      targetLevels.value = [5, 3, 1, 1]
      break
  }
}

// 获取概率颜色类
const getProbabilityClass = () => {
  if (!result.value) return ''
  const percent = result.value.probabilityPercent
  if (percent >= 75) return 'glow-text success'
  if (percent >= 50) return 'glow-text'
  if (percent >= 25) return 'warning'
  return 'danger'
}

// 获取概率提示
const getProbabilityHint = () => {
  if (!result.value) return ''
  const percent = result.value.probabilityPercent
  if (percent >= 75) return '成功率很高，祝你好运！'
  if (percent >= 50) return '成功率适中，值得一试'
  if (percent >= 25) return '成功率较低，请谨慎考虑'
  if (percent >= 10) return '成功率很低，建议调整目标'
  return '成功率极低，不建议尝试'
}

// 获取标签类型
const getProbabilityTagType = () => {
  if (!result.value) return 'info'
  const percent = result.value.probabilityPercent
  if (percent >= 75) return 'success'
  if (percent >= 50) return ''
  if (percent >= 25) return 'warning'
  return 'danger'
}
</script>

<style lang="scss" scoped>
@use 'sass:color';
@use '@/styles/variables' as *;

.strengthen-probability-page {
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
  
  .level-inputs {
    display: grid;
    grid-template-columns: repeat(4, 1fr);
    gap: $spacing-sm;
    
    :deep(.el-input-number) {
      width: 100%;
    }
  }
  
  .preset-buttons {
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
      margin-bottom: $spacing-lg;
      
      .probability-value {
        font-family: $font-tech;
        font-size: 4rem;
        font-weight: 900;
        line-height: 1;
        transition: color $transition-normal;
        
        &.success {
          color: $success-color;
        }
        
        &.warning {
          color: $warning-color;
        }
        
        &.danger {
          color: $danger-color;
        }
      }
      
      .probability-label {
        color: $text-secondary;
        margin-top: $spacing-sm;
      }
    }
    
    .probability-hint {
      text-align: center;
      margin-bottom: $spacing-xl;
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
    
    .paths-section {
      margin-top: $spacing-xl;
      
      .subsection-title {
        font-size: 1.2rem;
        margin-bottom: $spacing-md;
        color: $text-primary;
      }
      
      .paths-list {
        max-height: 300px;
        overflow-y: auto;
        padding: $spacing-md;
        background: rgba(0, 0, 0, 0.3);
        border: 1px solid $border-color;
        border-radius: $radius-md;
        
        .path-item {
          display: flex;
          gap: $spacing-md;
          padding: $spacing-sm;
          margin-bottom: $spacing-sm;
          font-family: $font-tech;
          font-size: 0.9rem;
          
          &.success {
            color: $success-color;
          }
          
          .path-index {
            color: $text-muted;
            min-width: 40px;
          }
          
          .path-result {
            min-width: 60px;
          }
        }
      }
    }
    
    .visualization-section {
      margin-top: $spacing-xl;
      
      .subsection-title {
        font-size: 1.2rem;
        margin-bottom: $spacing-md;
        color: $text-primary;
      }
      
      .bar-chart {
        height: 60px;
        background: rgba(0, 0, 0, 0.3);
        border: 1px solid $border-color;
        border-radius: $radius-sm;
        overflow: hidden;
        display: flex;
        
        .bar {
          height: 100%;
          display: flex;
          align-items: center;
          justify-content: center;
          color: $text-primary;
          font-weight: 500;
          transition: width $transition-slow ease-out;
          
          &.success-bar {
            background: linear-gradient(90deg, $success-color, color.scale($success-color, $lightness: 10%));
          }
          
          &.fail-bar {
            background: linear-gradient(90deg, $danger-color, color.scale($danger-color, $lightness: 10%));
          }
        }
      }
    }
  }
}

// Element Plus 样式覆盖
:deep(.el-input-number) {
  .el-input__inner {
    background: rgba(0, 0, 0, 0.5);
    border-color: $border-color;
    color: $text-primary;
    text-align: center;
    
    &:focus {
      border-color: $primary-color;
    }
  }
}

:deep(.el-radio) {
  .el-radio__label {
    color: $text-primary;
  }
  
  .el-radio__input.is-checked {
    .el-radio__inner {
      background: $primary-color;
      border-color: $primary-color;
    }
  }
}
</style>