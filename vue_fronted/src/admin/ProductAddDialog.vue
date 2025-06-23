<template>
  <el-dialog
      v-model="dialogVisible"
      title="新增商品"
      :close-on-click-modal="false"
      @close="handleClose"
      width="500px"
  >
    <el-form
        ref="formRef"
        :model="formData"
        :rules="formRules"
        label-width="80px"
    >
      <el-form-item label="商品名称" prop="name">
        <el-input v-model="formData.name" placeholder="请输入商品名称"></el-input>
      </el-form-item>

      <!-- 其他表单字段 -->
      <el-form-item label="商品价格" prop="price">
        <el-input v-model.number="formData.price" type="number" placeholder="请输入价格"></el-input>
      </el-form-item>
      <el-form-item label="商品数量" prop="num">
        <el-input v-model.number="formData.num" type="number" placeholder="请输入数量"></el-input>
      </el-form-item>
      <el-form-item label="商品单位" prop="unit">
        <el-input v-model.number="formData.unit" type="text" placeholder="请输入单位"></el-input>
      </el-form-item>
      <el-form-item label="商品描述" prop="desc">
        <el-input v-model.number="formData.unit" type="text" placeholder="请输入商品描述"></el-input>
      </el-form-item>
      <el-form-item label="商品图片" prop="pic" class="pic-select">
        <el-input v-model.number="formData.pic" type="button" placeholder="请输入商品描述"></el-input>
        <el-button type="primary" @click="resetForm">上传图片</el-button>
      </el-form-item>

      <!-- 省略其他表单字段（数量、单位、描述等） -->
    </el-form>

    <template #footer>
      <el-button @click="dialogVisible = false">取消</el-button>
      <el-button type="primary" @click="handleSubmit">确定</el-button>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, reactive, defineEmits, defineProps, watch } from 'vue';
import {ElMessage, ElForm, type FormInstance} from 'element-plus';
import axios from 'axios';

// 定义props和emits
const props = defineProps({
  visible: {
    type: Boolean,
    default: false
  }
});

let dialogVisible = ref(props.visible);

// 监听 props.visible 变化，同步到本地 ref
watch(() => props.visible, (newVal) => {
  dialogVisible.value = newVal;
});

const emits = defineEmits(['close', 'success']);

// 表单数据和验证规则




const formRef = ref<FormInstance>();
const formData = reactive<{
  name: '',
  price: 0,
  num: 0,
  pic:''
  unit: '',
  desc: ''
}>({
  name: '',
  price: 0,
  num: 0,
  pic: '',
  unit: '',
  desc: ''
});

const ALLOWED_UNITS = ['件', '盒', '瓶', '千克', '米'];

const formRules = reactive({
  name: [{ required: true, message: '请输入商品名称', trigger: 'blur' }],
  price: [{ required: true, message: '请输入价格', trigger: 'blur' }],
  num: [{ required: true, message: '请输入数量', trigger: 'blur' }],
  pic: [{ required: true, message: '请上传商品图片', trigger: 'blur' }],
  unit: [
    { required: true, message: '请输入商品单位', trigger: 'blur' },
    {
      validator: (_:any, value:any, callback:any) => {
        if (!value) return callback(); // 允许空值（由必填规则处理）
        if (!ALLOWED_UNITS.includes(value)) {
          callback(new Error(`单位只能是：${ALLOWED_UNITS.join('、')}`));
        } else {
          callback();
        }
      },
      trigger: ['blur', 'change'] // 失去焦点或值变化时验证
    }
  ],
  desc: [{ required: true, message: '请输入商品描述', trigger: 'blur' }],

  // 其他验证规则
});

// 提交表单
const handleSubmit = async () => {
  // await formRef.value?.validate((valid) => {
  //   if (!valid) return false;
  // });

  try {
    const response = await axios.post('http://localhost:8082/products',{
      name: formData.name,
      price: formData.price,
      num: formData.num,
      pic: formData.pic,
      unit: formData.unit,
      desc: formData.desc,
      create_time: new Date().toISOString() // 添加创建时间
    }, {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('admin_token')}` // 假设使用JWT令牌
      }
    })

    if (response.data.code === 200) {
      ElMessage.success('添加成功');
      emits('success'); // 通知父组件添加成功
      resetForm();
    } else {
      ElMessage.error(response.data.message || '添加失败');
    }
  } catch (error) {
    ElMessage.error('网络错误，请重试');
  }
};

// 重置表单
const resetForm = () => {
  formData.name = '';
  formData.price = 0;
  // 重置其他字段
  formRef.value?.resetFields();
};

// 关闭对话框
const handleClose = () => {
  resetForm();
  emits('close'); // 通知父组件对话框已关闭
};

// 监听props.visible变化
watch(() => props.visible, (newVal) => {
  if (!newVal) resetForm();
});
</script>

<style scoped lang="scss">
.pic-select {
  display: flex;
  align-items: center;

  .el-input {
    flex: 1;
    margin-right: 10px;
  }

  .el-button {
    flex-shrink: 0;
  }
}
</style>