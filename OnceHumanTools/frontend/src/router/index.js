import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('@/views/Home.vue'),
    meta: {
      title: '首页'
    }
  },
  {
    path: '/tools',
    name: 'Tools',
    component: () => import('@/views/Tools.vue'),
    children: [
      {
        path: 'affix-probability',
        name: 'AffixProbability',
        component: () => import('@/views/tools/AffixProbability.vue'),
        meta: {
          title: '模组词条概率计算器'
        }
      },
      {
        path: 'strengthen-probability',
        name: 'StrengthenProbability',
        component: () => import('@/views/tools/StrengthenProbability.vue'),
        meta: {
          title: '模组强化概率计算器'
        }
      }
    ]
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('@/views/About.vue'),
    meta: {
      title: '关于'
    }
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('@/views/NotFound.vue'),
    meta: {
      title: '页面未找到'
    }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { top: 0 }
    }
  }
})

// 路由守卫
router.beforeEach((to, from, next) => {
  // 设置页面标题
  document.title = to.meta.title ? `${to.meta.title} - OnceHuman工具集` : 'OnceHuman工具集'
  next()
})

export default router