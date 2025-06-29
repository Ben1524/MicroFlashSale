<template>
  <div class="common-layout">
    <div class="animated-background left"></div>
    <div class="centered-login-container">
      <el-form
          ref="loginRuleFormRef"
          :model="loginForm"
          :rules="rules"
          :label-width="formSize === 'default' ? '80px' : '60px'"
          :size="formSize"
          status-icon
      >
        <el-main>
          <el-form-item
              label="邮箱"
              prop="email"
              :label-width="formSize === 'default' ? '80px' : '60px'"
          >
            <el-input
                v-model="loginForm.email"
                placeholder="请输入邮箱"
                clearable
            ></el-input>
          </el-form-item>
          <el-form-item
              label="密码"
              prop="password"
              :label-width="formSize === 'default' ? '80px' : '60px'"
          >
            <el-input
                v-model="loginForm.password"
                type="password"
                placeholder="请输入密码"
                clearable
            ></el-input>
          </el-form-item>
        </el-main>
        <el-footer>
          <el-form-item>
            <el-button type="primary" @click="submitForm(loginRuleFormRef)">
              登录
            </el-button>
            <el-button @click="resetForm(loginRuleFormRef)">
              重置
            </el-button>
            <el-button type="text" @click="$router.push('/register')" class="login-link">
            没有账号？注册
            </el-button>
          </el-form-item>
        </el-footer>
      </el-form>
    </div>
    <div class="animated-background right"></div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue';
import type { ComponentSize, FormInstance, FormRules } from 'element-plus';
import apiClient from '@/services/api';
import { inject } from 'vue';
import { useRouter } from 'vue-router';

interface LoginRuleForm {
  email: string;
  password: string;
}

// 注入全局对象
const axios = inject('axios');
import auth from '@/auth/auth'
const router = useRouter();
const formSize = ref<ComponentSize>('large');
const loginRuleFormRef = ref<FormInstance>(); // 表单实例
const loginForm = reactive<LoginRuleForm>({ // 表单数据
  email: '',
  password: '',
});

const rules = reactive<FormRules<LoginRuleForm>>({ // 表单的验证规则
  email: [
    { required: true, message: 'Please input your email', trigger: 'blur' },
    { type: 'email', message: 'Please input a valid email address', trigger: ['blur', 'change'] },
  ],
  password: [
    { required: true, message: 'Please input your password', trigger: 'blur' },
    { min: 6, max: 20, message: 'Length should be 6 to 20', trigger: 'blur' },
    { pattern: /^[a-zA-Z0-9_]+$/, message: 'Password can only contain letters, numbers and underscores', trigger: 'blur' },
  ],
});

const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate(async (valid, fields) => {
    if (valid) {
      try {
        const response = await axios.post('http://localhost:8082/user/front_user_login', {
          email: loginForm.email,
          password: loginForm.password,
        });
        if (response.data.success) {
          alert(response.data.message);
          console.log(response.data)
          // 保存jwt-token
          auth.setFrontAuth(response.data.token,loginForm.email)
          router.push("/")
        } else {
          alert(response.data.message);
        }
      } catch (error) {
        alert('Network error');
      }
    } else {
      alert('Form validation failed');
    }
  });
};

const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.resetFields();
};

const options = Array.from({ length: 10000 }).map((_, idx) => ({
  value: `${idx + 1}`,
  label: `${idx + 1}`,
}));
</script>

<style scoped lang="scss">
.common-layout {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  position: relative;
  overflow: hidden;
}

.centered-login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1;
  width: 100%;
  height: 100%;
  max-width: 400px;
}

.animated-background {
  position: absolute;
  top: 0;
  width: 20%;
  height: 100%;
  background: linear-gradient(45deg, #3498db, #e74c3c);
  animation: moveBackground 10s infinite alternate;
  opacity: 0.5;
}

.left {
  left: 0;
}

.right {
  right: 0;
}

@keyframes moveBackground {
  from {
    transform: translateY(-100%);
  }
  to {
    transform: translateY(100%);
  }
}
</style>