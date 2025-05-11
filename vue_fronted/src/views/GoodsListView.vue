<template>
  <div class="page-container">
    <el-row :gutter="16">
      <el-col :span="6" v-for="seckill in seckills" :key="seckill.id">
        <router-link :to="{ name: 'GoodsDetail', params: { id: seckill.id } }" class="inner-box">
          <div class="thumbnail">
            <img :src="'data:image/jpeg;base64,' + seckill.pic" alt="">
          </div>
          <div class="goods-info">
            <div class="title"><b>{{ seckill.name }}</b></div>
            <div class="price">
              <b>秒杀价：￥{{ seckill.price }}</b> &nbsp;
              <del>原价:￥{{ seckill.p_price }}</del>
            </div>
            <div class="desc">{{ seckill.pdesc }}</div>
          </div>
        </router-link>
      </el-col>
    </el-row>
    <el-button @click="loadData" class="load_more">加载更多</el-button>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { ElMessage } from 'element-plus';
import { useRouter } from 'vue-router';
import axios from "axios";

const router = useRouter();
const currentPage = ref(1);
const pageSize = ref(8);
const seckills = ref([]);
const totalPage = ref(undefined);

const getSeckillList = async (currentPage, pageSize) => {
  try {
    const rep = await axios.get('/seckill/front/get_seckill_list', {
      params: {
        currentPage,
        pageSize
      }
    });
    seckills.value = seckills.value.concat(rep.data.seckill_list);
    currentPage.value = rep.data.current;
    pageSize.value = rep.data.page_size;
    totalPage.value = rep.data.total_page;
  } catch (err) {
    console.log(err);
    ElMessage.error('获取数据失败');
  }
};

const loadData = () => {
  if (totalPage.value >= (currentPage.value + 1)) {
    getSeckillList(currentPage.value + 1, pageSize.value);
  } else {
    ElMessage.warning('无更多数据');
  }
};

const logout = () => {
  // $auth.delFrontAuth();
  // router.push('/login');
};

onMounted(() => {
  getSeckillList(currentPage.value, pageSize.value);
});
</script>

<style scoped lang="scss">
.load_more {
  text-align: center;
  margin-top: 30px;
  background-color: rgb(194, 212, 235);
}


.inner-box {
  .thumbnail {
    img {
      width: 100%;
      height: 240px;
    }
  }
  .goods-info {
    padding: 10px;
    .price {
      color: brown;
      padding: 10px 0;
      text-align: center;
      font-size: 16px;
      del {
        text-decoration: line-through;
        font-size: 14px;
      }
    }
    .desc {
      color: #666;
    }
  }
}
</style>