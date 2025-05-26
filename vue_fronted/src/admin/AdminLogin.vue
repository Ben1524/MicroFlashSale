<template>
  <div class="admin-layout">
    <div class="admin-background"></div>
    <div class="admin-login-container">
      <h1 class="admin-title">管理员登录</h1>
      <p class="admin-tip">仅限系统管理员访问</p>

      <el-form
          ref="loginFormRef"
          :model="loginForm"
          :rules="rules"
          :label-width="formSize === 'default' ? '80px' : '60px'"
          :size="formSize"
          status-icon
      >
        <el-form-item label="用户名" prop="username" :label-width="'80px'">
          <el-input
              v-model="loginForm.username"
              placeholder="请输入管理员用户名"
              clearable
          ></el-input>
        </el-form-item>

        <el-form-item label="密码" prop="password" :label-width="'80px'">
          <el-input
              v-model="loginForm.password"
              type="password"
              placeholder="请输入管理员密码"
              clearable
          ></el-input>
        </el-form-item>

        <el-form-item class="button-group">
          <el-button type="primary" @click="handleLogin">
            登录后台
          </el-button>
          <el-button @click="resetForm">
            重置
          </el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue';
import type { FormInstance, FormRules } from 'element-plus';
import { useRouter } from 'vue-router';
import { inject } from 'vue';
import auth from '@/auth/auth'; // 假设存在管理员专用认证模块

// 注入HTTP客户端
const axios = inject('axios') as any;
const router = useRouter();

// 表单状态
const formSize = ref<'large' | 'default'>('large');
const loginFormRef = ref<FormInstance>();
const loginForm = reactive<{
  username: string;
  password: string;
}>({
  username: '',
  password: '',
});

// 验证规则
const rules = reactive<FormRules<typeof loginForm>>({
  username: [
    { required: true, message: '请输入管理员用户名', trigger: 'blur' },
  ],
  password: [
    { required: true, message: '请输入管理员密码', trigger: 'blur' },
    { min:6 , max: 20, message: '密码长度需6-20位', trigger: 'blur' },
    { pattern: /^[A-Za-z0-9_@!#$%^&*]+$/, message: '密码包含字母、数字和特殊符号', trigger: 'blur' },
  ],
});

// 登录处理
const handleLogin = async () => {
  const form = loginFormRef.value;
  if (!form) return;

  await form.validate(valid => {
    if (valid) {
      axios.post('http://localhost:8082/user/admin_login', {
        username: loginForm.username,
        password: loginForm.password
      }).then((res: any) => {
        if (res.data.success) {
          // 保存管理员Token（假设auth模块有专门方法）
          auth.setAdminToken(res.data.token);
          // 跳转至管理员后台
          router.push('/admin/dashboard');
          alert('登录成功，欢迎进入管理后台');
        } else {
          alert(`登录失败：${res.data.message}`);
        }
      }).catch(() => {
        alert('网络请求失败，请检查连接');
      });
    }
  });
};

// 重置表单
const resetForm = () => {
  loginFormRef.value?.resetFields();
};
</script>

<style scoped lang="scss">
.admin-layout {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f0f2f5;
  position: relative;
}

.admin-background {
  position: absolute;
  width: 100%;
  height: 100%;
  background: linear-gradient(135deg, #1a73e8, #1553a1);
  opacity: 0.05;
  filter: blur(100px);
}

.admin-login-container {
  background: white;
  padding: 3rem 2.5rem;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0,0,0,0.08);
  width: 100%;
  max-width: 400px;
}

.admin-title {
  color: #1a73e8;
  font-size: 1.8rem;
  font-weight: 600;
  text-align: center;
  margin-bottom: 1.2rem;
}

.admin-tip {
  color: #666;
  font-size: 0.9rem;
  text-align: center;
  margin-bottom: 2rem;
  opacity: 0.8;
}

.button-group {
  display: flex;
  justify-content: space-between;
  margin-top: 2rem;
}

.el-form-item__label {
  color: #333;
  font-weight: 500;
}

.el-input {
  border-radius: 8px;
}

.el-button--primary {
  background-color: #1a73e8;
  border-color: #1a73e8;
  width: 48%;
}

.el-button--primary:hover {
  background-color: #1553a1;
  border-color: #1553a1;
}
</style>