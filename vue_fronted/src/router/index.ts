import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'Home',
      component: () => import('../views/HomeView.vue')
    },
    {
      path: '/detail/:id',
      name: 'GoodsDetail',
      component: () => import('../views/GoodsDetailView.vue')
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/LoginView.vue')
    },
    {
      path: '/register',
      name: 'register',
      component: () => import('../views/RegisterView.vue')
    },
    // 管理员路由
    {
      path: '/admin',
      name: 'Admin',
      component: () => import('../admin/AdminLayout.vue'), // 新增布局组件
      meta: {requiresAdmin: true},
      children: [
        {
          path: 'dashboard', // 默认子路由
          name: 'AdminDashboard',
          component: () => import('../admin/DashBoard.vue'),
          children: [
            {
              path: 'users',
              name: 'UserList',
              component: () => import('../admin/UserList.vue')
            },
            {
              path: 'products',
                name: 'ProductList',
                component: () => import('../admin/ProductList.vue')
            }
              ]
        },
        {
          path: 'login',
          name: 'AdminLogin',
          component: () => import('../admin/AdminLogin.vue')
        },
      ]
    }
  ]
})

router.beforeEach((to, from, next) => {
  // 公共路径直接放行
  if (to.path === "/" || to.path === "/login" || to.path === "/register" || to.path === "/admin") {
    next()
  } else {
    // 使用 startsWith 替代正则，更简洁准确
    if (to.path.startsWith('/admin/')) {
      // 管理员路径校验
      const admin_token = localStorage.getItem("admin_token")

      if (!admin_token) {
        next("/admin") // 无token跳转到管理员登录页
      } else {
        next() // 有token放行
      }
    } else {
      // 用户路径校验
      const front_token = localStorage.getItem("front_token")

      if (!front_token) {
        next("/login")
      } else {
        next()
      }
    }
  }
})

export default router