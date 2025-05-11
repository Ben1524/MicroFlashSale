<template>
  <div class="common-layout">
    <div class="animated-background left"></div>
    <div class="centered-login-container">
      <el-form
          ref="registerRuleFormRef"
          :model="registerForm"
          :rules="rules"
          :label-width="formSize === 'default' ? '100px' : '80px'"
          :size="formSize"
          status-icon
      >
        <el-form-item
            label="邮箱"
            prop="email"
            :label-width="formSize === 'default' ? '100px' : '80px'"
        >
          <el-input
              v-model="registerForm.email"
              placeholder="请输入邮箱"
              clearable
              prefix-icon="el-icon-mail"
          ></el-input>
        </el-form-item>
        <el-form-item
            label="验证码"
            prop="code"
            :label-width="formSize === 'default' ? '100px' : '80px'"
        >
          <div class="input-code-container"> <!-- 新增容器用于水平布局 -->
            <el-input
                v-model="registerForm.code"
                placeholder="请输入邮箱验证码"
                clearable
                style="width: 200px;"
            ></el-input>
            <el-button
                class="verification-code"
                type="primary"
                @click="sendVerificationCode"
            >获取验证码</el-button>
          </div>
        </el-form-item>

        <el-form-item
            label="密码"
            prop="password"
            :label-width="formSize === 'default' ? '100px' : '80px'"
        >
          <el-input
              v-model="registerForm.password"
              type="password"
              placeholder="请输入密码"
              clearable
              prefix-icon="el-icon-lock"
          ></el-input>
        </el-form-item>

        <el-form-item
            label="确认密码"
            prop="confirmPassword"
            :label-width="'80px'"
        >
          <el-input
              v-model="registerForm.confirmPassword"
              type="password"
              placeholder="请再次输入密码"
              clearable
              prefix-icon="el-icon-lock"
          ></el-input>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="submitForm(registerRuleFormRef)">
            注册
          </el-button>
          <el-button @click="resetForm(registerRuleFormRef)">
            重置
          </el-button>
          <el-button type="text" @click="$router.push('/login')" class="login-link">
            已有账号？登录
          </el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="animated-background right"></div>
  </div>
</template>


<script setup lang="ts">
import { reactive, ref } from 'vue';
import type { ComponentSize, FormInstance, FormRules } from 'element-plus';
import axios from 'axios';
import { useRouter } from 'vue-router';

interface RegisterRuleForm {
  email: string;
  password: string;
  code:string
  confirmPassword: string;
}

const router = useRouter();
const formSize = ref<ComponentSize>('large');
const registerRuleFormRef = ref<FormInstance>();
const registerForm = reactive<RegisterRuleForm>({
  email: '',
  password: '',
  code:'',
  confirmPassword: '',
});

// 自定义密码一致性验证
const validatePassword = (rule: any, value: string, callback: any) => {
  if (value.length < 6 || value.length > 20) {
    callback(new Error('密码长度应为6-20位'));
  } else if (!/^[a-zA-Z0-9_]+$/.test(value)) {
    callback(new Error('密码只能包含字母、数字和下划线'));
  } else {
    callback();
  }
};

const validateConfirmPassword = (rule: any, value: string, callback: any) => {
  if (value !== registerForm.password) {
    callback(new Error('两次输入的密码不一致'));
  } else {
    callback();
  }
};

const rules = reactive<FormRules<RegisterRuleForm>>({
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入有效的邮箱地址', trigger: ['blur', 'change'] },
  ],
  code: [
    { required: true, message: '请输入验证码', trigger: 'blur' },
    { min: 6, max: 6, message: '验证码长度应为6位', trigger: 'blur' },
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { validator: validatePassword, trigger: 'blur' },
  ],
  confirmPassword: [
    { required: true, message: '请再次输入密码', trigger: 'blur' },
    { validator: validateConfirmPassword, trigger: 'blur' },
  ],
});

const submitForm = async (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  await formEl.validate(async (valid, fields) => {
    if (valid) {
      try {
        const response = await axios.post('/register', registerForm);
        if (response.data.success) {
          alert('注册成功！请登录');
          router.push('/login');
        } else {
          alert(response.data.message);
        }
      } catch (error) {
        alert('注册失败，请检查网络或重试');
      }
    } else {
      alert('表单验证失败，请检查输入内容');
    }
  });
};

const sendVerificationCode = async () => {
  if (!registerForm.email) {
    alert('请检查邮箱');
    return;
  }
  try {
    console.log(registerForm.email);
    const response = await axios.post('http://localhost:8082/user/send_email', { email: registerForm.email });

    if (response.data.success) {
      alert('验证码已发送，请检查邮箱');
    } else {
      alert(response.data.message);
    }
  } catch (error) {
    alert('发送验证码失败，请检查网络或重试');
  }
};

const resetForm = (formEl: FormInstance | undefined) => {
  if (!formEl) return;
  formEl.resetFields();
};
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
// 新增水平布局样式
.input-code-container {
  display: flex;
  align-items: center;
  //gap: 10px; // 输入框和按钮之间的间距
}

// 原有样式保持不变，可根据需要调整输入框宽度
.el-input {
  // 如需统一调整输入框样式可在此添加
}
@keyframes moveBackground {
  from {
    transform: translateY(-100%);
  }
  to {
    transform: translateY(100%);
  }
}

.login-link {
  float: right;
  color: #409eff;
  font-size: 14px;
}
</style>