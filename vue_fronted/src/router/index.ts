import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'

const router = createRouter({
    history:createWebHistory(),
    routes:[
  {
    path: '/',
    name: 'Home',
    component: () => import('../views/HomeView.vue')
  },
  {
    path: '/detail/:id',
    name: 'GoodsDetail',
    component: () => import(/* webpackChunkName: "about" */ '../views/GoodsDetailView.vue')
  },
  {
    path:'/login',
    name:'login',
    component: () => import('../views/LoginView.vue')
  },
  {
    path:'/register',
    name:'register',
    component: () => import('../views/RegisterView.vue')
  }
]})

// 路由守卫
router.beforeEach((to,from,next) => {
  // to 满足条件的url放过
  if (to.path == "/" || to.path == "/login" || to.path == "/register" || to.path == "/admin"){
    next()  // 直接放过
  }else {
    const reg = /\/main\/[a-zA-z]+[_-]?/
    // 做token校验
    if (reg.test(to.path)) {
      // 校验管理员，不受用户影响
      const admin_token = localStorage.getItem("admin_token")
      // alert("到main的url")
      if (admin_token == "" || admin_token == null || admin_token == '') {
        next("/admin")
      }else {
        next()
      }
    }else {
      // 校验用户的，不受管理员影响
      const front_token = localStorage.getItem("front_token")
      // front的需要权限的url
      if (front_token == "" || front_token == null || front_token == '') {
        next("/login")
      }else {
        next()
      }
    }
  }
})

export default router
