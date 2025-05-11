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
import axios from 'axios';

interface LoginRuleForm {
  email: string;
  password: string;
}

const formSize = ref<ComponentSize>('large');
const loginRuleFormRef = ref<FormInstance>();
const loginForm = reactive<LoginRuleForm>({
  email: '',
  password: '',
});

const rules = reactive<FormRules<LoginRuleForm>>({
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
        const response = await axios.post('/login', loginForm);
        if (response.data.success) {
          alert(response.data.message);
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