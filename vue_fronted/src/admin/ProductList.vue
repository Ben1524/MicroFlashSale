<template>
  <div class="product-list-container">
    <!-- 顶部搜索区域 -->
    <div class="search-bar mb-20">
      <el-row :gutter="14" justify="start">
        <el-col :span="2">
          <el-button type="primary" @click="addProducts">
            添加商品
          </el-button>
        </el-col>
        <el-col :span="6">
          <el-input
              v-model="searchQuery"
              placeholder="搜索邮箱"
              prefix-icon="el-icon-search"
              @keyup.enter="fetchProducts"
          ></el-input>
        </el-col>
        <el-col :span="6">
          <el-select
              v-model="currentStatus"
              placeholder="筛选状态"
              @change="fetchProducts"
          >
            <el-option label="全部" value="-1"></el-option>
            <el-option label="启用" value="1"></el-option>
            <el-option label="禁用" value="0"></el-option>
          </el-select>
        </el-col>
      </el-row>
    </div>

    <!-- 表格内容区域 -->
    <div class="content-wrapper">
      <el-table
          :data="users"
          stripe
          border
          v-loading="isLoading"
          element-loading-text="数据加载中..."
          @sort-change="handleSortChange"
      >
        <el-table-column type="index" width="50"></el-table-column>
        <el-table-column prop="id" label="ID" width="80"></el-table-column>
        <el-table-column prop="name" label="名称" show-overflow-tooltip></el-table-column>
        <el-table-column prop="price" label="价格" show-overflow-tooltip></el-table-column>
        <el-table-column prop="num" label="数量" width="180"></el-table-column>
        <el-table-column prop="unit" label="单位" width="100"></el-table-column>
        <el-table-column prop="pic" label="图片" width="120"></el-table-column>
        <el-table-column prop="desc" label="描述" show-overflow-tooltip></el-table-column>
        <el-table-column prop="create_time" label="创建时间" width="180">
          <template #default="{ row }">
            {{ formatDate(row.create_time) }}
          </template>
        </el-table-column>
      </el-table>
    </div>

    <ProductAddDialog
        :visible="dialogVisible"
        @close="dialogVisible = false"
        @success="fetchProducts"
/>


    <!-- 分页控件 -->
    <div class="pagination-container">
      <el-pagination
          :current-page="currentPage"
          :page-sizes="[10, 20, 50, 100]"
          :page-size="pageSize"
          :total="total"
          @size-change="handlePageSizeChange"
          @current-change="handleCurrentPageChange"
          layout="total, sizes, prev, pager, next, jumper"
      ></el-pagination>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue';
import { ElMessage } from 'element-plus';
import axios from "axios";
import router from "@/router";
import ProductAddDialog from "@/admin/ProductAddDialog.vue";

// 状态管理
const searchQuery = ref('');
const currentStatus = ref(-1); // -1表示全部，0=禁用，1=启用
const currentPage = ref(1);
const pageSize = ref(10);
const total = ref(0);
const users = ref<any[]>([]);
const isLoading = ref(false);
const sortField = ref('');
const sortOrder = ref('');
let dialogVisible = ref(false);

// 日期格式化
const formatDate = (dateString: string) => {
  if (!dateString) return '-';
  const date = new Date(dateString);
  return date.toISOString().split('T')[0];
};

// 生成测试数据（模拟后端返回格式）
const generateTestData = (totalCount = 100) => {
  const data = [];
  for (let i = 0; i < totalCount; i++) {
    data.push({
      id: i + 1,
      email: `user${i + 1}@example.com`,
      desc: `用户类型：${i % 2 === 0 ? '普通用户' : 'VIP用户'}`,
      status: Math.floor(Math.random() * 2), // 随机生成0或1
      create_time: new Date(Date.now() - i * 24 * 60 * 60 * 1000).toISOString()
    });
  }
  return data;
};

// 模拟分页数据
const getPagedData = (data: any[], page: number, size: number) => {
  const start = (page - 1) * size;
  const end = page * size;
  return data.slice(start, end);
};

// 获取用户列表（使用测试数据）
const fetchProducts = async () => {


  let curP = currentPage.value;
  let size = pageSize.value;

  // 向/user/users 接口请求数据
  try {
    const response = await axios.get('http://localhost:8082/product/products', {
      params: {
        CurrentPage: currentPage.value,
        PageSize: pageSize.value,
        Search: searchQuery.value,
        Status: currentStatus.value,
        SortField: sortField.value,
        SortOrder: sortOrder.value
      },
      headers:{
        Authorization: `Bearer ${localStorage.getItem('admin_token')}` // 假设使用JWT令牌,bearer标识是可选的
      }
    });

    if (!response.data||response.status !== 200) {
      await router.push('/admin/login'); // 如果状态码不是200，重定向到登录页面
    }

    total.value = response.data.total;
    users.value = response.data.products;
    console.log('成功获取用户列表', users.value);

  } catch (error:any) {
    let message = '获取用户列表失败';
    if (error.response) {
      message = error.response.data.message || message;
    } else if (error.request) {
      message = '请求发送失败，未收到响应';
    }
    ElMessage.error(message);
    console.error('请求详情', error.config.params); // 打印最终发送的参数
  }
  return;




  //
  // isLoading.value = true;
  //
  // // 生成模拟数据（可根据筛选条件过滤）
  // const allData = generateTestData(100);
  // let filteredData = allData;
  //
  // // 搜索过滤
  // if (searchQuery.value) {
  //   filteredData = filteredData.filter(item =>
  //       item.email.includes(searchQuery.value)
  //   );
  // }
  //
  // // 状态过滤
  // if (currentStatus.value !== -1) {
  //   filteredData = filteredData.filter(item =>
  //       item.status === currentStatus.value
  //   );
  // }
  //
  // // 排序处理
  // if (sortField.value) {
  //   filteredData = filteredData.sort((a :any, b:any) => {
  //     const order = sortOrder.value === 'asc' ? 1 : -1;
  //     if (a[sortField.value] < b[sortField.value]) return -1 * order;
  //     if (a[sortField.value] > b[sortField.value]) return 1 * order;
  //     return 0;
  //   });
  // }
  //
  // // 设置分页数据
  // total.value = filteredData.length;
  // users.value = getPagedData(filteredData, currentPage.value, pageSize.value);
  //
  // isLoading.value = false;
};

// 其他函数保持不变
const handleSortChange = ({ prop, order }: { prop: string; order: string }) => {
  sortField.value = prop;
  sortOrder.value = order === 'ascending' ? 'asc' : 'desc';
  fetchProducts();
};


const addProducts = () => {
  // router.push('/admin/product/add');
  dialogVisible.value = true; // 打开添加商品对话框
  console.log('打开添加商品对话框');
};


const handlePageSizeChange = (size: number) => {
  pageSize.value = size;
  currentPage.value = 1;
  fetchProducts();
};

const handleCurrentPageChange = (page: number) => {
  currentPage.value = page;
  fetchProducts();
};

// 初始化加载
onMounted(() => {
  fetchProducts();
});
</script>


<style scoped>
.product-list-container {
  display: flex;
  flex-direction: column;
  height: 70vh;
  padding: 30px;
  background-color: #fff;
}

.search-bar {
  padding-bottom: 10px;
  border-bottom: 1px solid #ebeef5;
}

.content-wrapper {
  flex: 1;
  min-height: 0; /* 允许内容区域收缩 */
  margin-bottom: 20px;
  overflow-y: auto; /* 内容过多时显示滚动条 */
}

.pagination-container {
  display: flex;
  justify-content: center;
  padding: 5px 0;
  background-color: #fff;
  border-top: 1px solid #ebeef5;
}

.mb-20 {
  margin-bottom: 20px;
}

/* 优化表格样式 */
.el-table {
  border-radius: 4px;
}

@media (max-width: 768px) {
  .search-bar .el-col {
    margin-bottom: 10px;
  }
}
</style>