<template>
  <div class="admin-dashboard">
    <el-container>
      <!-- 顶部导航栏 -->
      <el-header class="header">
        <div class="logo">管理系统</div>
        <div class="user-info">
          <!-- 头像作为触发区域 -->
          <el-dropdown @command="handleCommand" trigger="hover" placement="bottom-start">
            <div class="avatar-wrapper">
              <el-avatar
                  size="small"
                  src="https://picsum.photos/100/100"
                  class="hover-pointer"
              ></el-avatar>
<!--              &lt;!&ndash; 悬浮时显示下拉箭头（可选） &ndash;&gt;-->
<!--              <el-icon class="arrow-icon">-->
<!--                <ArrowDown />-->
<!--              </el-icon>-->
            </div>

            <el-dropdown-menu class="user-dropdown">
              <el-dropdown-item command="profile">个人信息</el-dropdown-item>
              <el-dropdown-item command="settings">设置</el-dropdown-item>
              <el-dropdown-item command="logout">退出登录</el-dropdown-item>
            </el-dropdown-menu>
          </el-dropdown>
          <span class="admin-name">管理员</span> <!-- 单独显示用户名 -->
        </div>
      </el-header>

      <!-- 主体内容区 -->
      <el-container>
        <!-- 侧边栏 -->
        <el-aside class="sidebar">
          <el-menu
              :default-active="activeMenu"
              class="el-menu-vertical"
              router
              background-color="#304156"
              text-color="#fff"
              active-text-color="#409EFF"
              @open="handleOpen"
              @close="handleClose">

            <el-sub-menu index="1">
              <template #title>
                <i class="el-icon-user"></i>
                <span>用户管理</span>
              </template>
              <el-menu-item index="/admin/dashboard/users">用户列表</el-menu-item>
            </el-sub-menu>

            <el-sub-menu index="2">
              <template #title>
                <i class="el-icon-shopping-bag-2"></i>
                <span>商品管理</span>
              </template>
              <el-menu-item index="/admin/dashboard/products">商品列表</el-menu-item>
              <el-menu-item index="/activity-management">活动管理</el-menu-item>
            </el-sub-menu>
          </el-menu>
        </el-aside>

        <!-- 主内容区 -->
        <el-main class="main-content">
          <el-breadcrumb separator-class="el-icon-arrow-right">
            <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
            <el-breadcrumb-item>{{ breadcrumbName }}</el-breadcrumb-item>
          </el-breadcrumb>

          <div class="content-container">
            <router-view></router-view>
          </div>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import {ArrowDown} from "@element-plus/icons-vue";

const route = useRoute();
const router = useRouter();

// 当前激活的菜单
const activeMenu = computed(() => route.path);

// 面包屑名称映射
const breadcrumbMap = {
  '/admin/dashboard/users': '用户列表',
  '/admin/dashboard/products': '商品列表',
  '/activity-management': '活动管理'
};

// 当前面包屑名称
const breadcrumbName = computed(() => {
  return breadcrumbMap[route.path] || '未知页面';
});

// 菜单展开/收起事件
const handleOpen = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
};

const handleClose = (key: string, keyPath: string[]) => {
  console.log(key, keyPath);
};

// 顶部导航栏下拉菜单命令处理
const handleCommand = (command: string) => {
  switch (command) {
    case 'logout':
      // 登出逻辑
      localStorage.removeItem('admin_token');
      router.push('/admin/login');
      break;
    case 'profile':
      // 个人信息
      router.push('/admin/profile');
      break;
    case 'settings':
      // 设置页面
      router.push('/admin/settings');
      break;
  }
};

// 初始化处理
onMounted(() => {
  // 可以添加页面初始化逻辑
});
</script>

<style scoped>
.admin-dashboard {
  height: 100vh;
}

.header {
  background-color: #409EFF;
  color: white;
  display: flex;
  align-items: center;
  justify-content: space-between;  /*使logo和用户信息两端对齐*/
  padding: 0 20px;
  overflow: visible !important;
  /* 确保头部不遮挡下拉菜单 */
  position: relative;
  z-index: 999;
}

.logo {
  font-size: 20px;
  font-weight: bold;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 10px;
}

.sidebar {
  min-width: 200px;
  background-color: #304156;
}

.main-content {
  padding: 20px;
  background-color: #f0f2f5;
}

.content-container {
  margin-top: 10px;
  background-color: white;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
  min-height: calc(100vh - 140px); /* 计算内容区高度，避免滚动条 */
}

/* 响应式布局 */
@media (max-width: 768px) {
  .sidebar {
    min-width: 64px; /* 移动端侧边栏收窄 */
  }

  .el-menu--vertical {
    width: 64px;
  }

  .el-menu--vertical .el-sub-menu__title span,
  .el-menu--vertical .el-menu-item span {
    display: none; /* 隐藏文字 */
  }

  .user-info {
    /* 确保移动端有足够空间显示菜单 */
    padding-right: 20px;
  }

  .el-dropdown-menu {
    /* 移动端强制居右显示 */
    right: 20px !important;
  }
}

.el-dropdown-menu {
  z-index: 3000; /* 确保高于其他元素 */
}
</style>