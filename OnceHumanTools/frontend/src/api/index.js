import axios from 'axios'
import { ElMessage } from 'element-plus'

// 创建axios实例
const request = axios.create({
  baseURL: '/api/v1',
  timeout: 30000,
  headers: {
    'Content-Type': 'application/json'
  }
})

// 请求拦截器
request.interceptors.request.use(
  config => {
    // 可以在这里添加token等认证信息
    return config
  },
  error => {
    console.error('请求错误:', error)
    return Promise.reject(error)
  }
)

// 响应拦截器
request.interceptors.response.use(
  response => {
    return response.data
  },
  error => {
    const message = error.response?.data?.message || '网络错误'
    ElMessage.error(message)
    return Promise.reject(error)
  }
)

// API接口定义
export const api = {
  // 系统接口
  system: {
    healthCheck: () => request.get('/health')
  },
  
  // 模组接口
  mod: {
    // 获取词条列表
    getAffixList: () => request.get('/mod/affix/list'),
    
    // 计算词条概率
    calculateAffixProbability: (data) => request.post('/mod/affix/probability', data),
    
    // 计算强化概率
    calculateStrengthenProbability: (data) => request.post('/mod/strengthen/probability', data)
  },
  
  // 工具接口
  tools: {
    // 获取工具列表
    getToolsList: () => request.get('/tools')
  }
}

export default api