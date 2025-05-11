<template>
  <div class="page-container">
    <el-page-header content="商品详情页"></el-page-header>
    <el-card class="main-body">
      <div class="goods-detail" style="display: flex; justify-content: flex-start;">
        <div class="picture-box">
          <img :src="'data:image/jpeg;base64,' + pic" alt="" style="width: 380px; height: 380px;">
        </div>
        <div class="goods-info" style="margin-left: 40px; flex: 1; text-align: left;">
          <el-card :header="name">
            <template #content>
              <p class="desc">{{ pdesc }}</p>
              <ul class="text-info">
                <li>
                  <span class="label">秒杀价</span>
                  <span class="seckill-price">¥{{ price }}</span>
                  <span class="real-price">¥{{ p_price }}</span>
                </li>
                <li>
                  <span class="label">库存</span>
                  <span class="value">{{ num }}</span>
                </li>
                <li>
                  <span class="label">单位</span>
                  <span class="value">{{ unit }}</span>
                </li>
                <li>
                  <span class="label">秒杀时间</span>
                  <span class="seckill-time">{{ start_time }} -- {{ end_time }}</span>
                </li>
                <li>
                  <span class="label">抢购倒计时</span>
                  <span class="seckill-time" style="color:red;font-size:18px;font-weight: 550;" v-if="message === ''">
                    {{ hours }}:{{ minutes }}:{{ seconds }}
                  </span>
                  <span class="seckill-time" style="color:red;font-size:18px;font-weight: 550;" v-else>
                    {{ message }}
                  </span>
                </li>
              </ul>
              <el-button type="primary" @click="onSubmit">立即购买</el-button>
            </template>
          </el-card>
        </div>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import qs from 'qs';
import { ElMessage } from 'element-plus';
import { useRoute } from 'vue-router';

const route = useRoute();
const id = ref(undefined);
const name = ref('');
const num = ref(undefined);
const price = ref(undefined);
const pid = ref(undefined);
const pic = ref('');
const p_price = ref('');
const pdesc = ref('');
const unit = ref('');
const start_time = ref('');
const end_time = ref('');
const hours = ref('');
const minutes = ref('');
const seconds = ref('');
const message = ref('');

const getResult = async () => {
  try {
    const rep = await axios.get('/seckill/front/get_seckill_result', {
      headers: {
        Authorization: `Bearer ${localStorage.getItem('front_token')}`
      }
    });
    console.log(rep.data);
    console.log(rep.data.code);

    if (rep.data.code === 200) {
      ElMessage.success(rep.data.msg);
    } else {
      setTimeout(() => {
        getResult();
      }, 2000);
    }
  } catch (err) {
    ElMessage.error(err.message);
  }
};

const onSubmit = () => {
  const startTime = Date.parse(start_time.value);
  const nowTime = Date.parse(new Date());
  const sub = startTime - nowTime;

  if (sub > 0) {
    ElMessage.warning('抢购还未开始');
    return;
  }

  axios.post('/seckill/front/seckill', qs.stringify({
    id: id.value
  }), {
    headers: {
      Authorization: `Bearer ${localStorage.getItem('front_token')}`
    }
  }).then((rep) => {
    if (rep.data.code === 200) {
      ElMessage.success(rep.data.msg);
      window.location.reload();
    } else {
      getResult();
      ElMessage.warning(rep.data.msg);
    }
  }).catch((err) => {
    console.log(err);
    ElMessage.error(err.message);
  });
};

const get_seckill_by_id = () => {
  axios.get('/seckill/front/seckill_detail', {
    params: {
      id: route.params.id
    },
    headers: {
      Authorization: `Bearer ${localStorage.getItem('front_token')}`
    }
  }).then((rep) => {
    console.log(rep);
    if (rep.data.code === 200) {
      id.value = rep.data.seckill.id;
      name.value = rep.data.seckill.name;
      num.value = rep.data.seckill.num;
      price.value = rep.data.seckill.price;
      pid.value = rep.data.seckill.pid;
      pic.value = rep.data.seckill.pic;
      p_price.value = rep.data.seckill.p_price;
      pdesc.value = rep.data.seckill.pdesc;
      unit.value = rep.data.seckill.unit;
      start_time.value = rep.data.seckill.start_time;
      end_time.value = rep.data.seckill.end_time;
    } else {
      ElMessage.warning(rep.data.msg);
      console.log(rep.data.msg);
    }
  }).catch((err) => {
    console.log(err);
    ElMessage.error(err.message);
  });
};

onMounted(() => {
  get_seckill_by_id();

  setInterval(() => {
    const startTime = Date.parse(start_time.value);
    const nowTime = Date.parse(new Date());
    const sub = startTime - nowTime;

    if (sub > 0) {
      const h = Math.floor((sub % (24 * 60 * 60 * 1000)) / (60 * 60 * 1000));
      const m = Math.floor((sub % (60 * 60 * 1000)) / (60 * 1000));
      const s = Math.floor((sub % (60 * 1000)) / (1000));

      hours.value = h;
      minutes.value = m;
      seconds.value = s;
    } else {
      message.value = '抢购已开始';
    }
  }, 1000);
});
</script>

<style scoped lang="scss">
.main-body {
  width: 100%;
  .goods-detail {
    .picture-box {
      img {
        width: 100%;
        height: 100%;
      }
    }
    .goods-info {
      .desc {
        padding: 5px 0;
        color: #666666;
        font-size: 13px;
      }
      .text-info {
        padding-top: 30px;
        li {
          display: flex;
          flex-direction: row;
          justify-content: flex-start;
          align-items: center;
          padding: 10px;
          .label {
            font-weight: 600;
            width: 100px;
          }
          .seckill-price {
            font-size: 26px;
            color: crimson;
            font-weight: 600;
          }
          .real-price {
            margin-left: 10px;
            font-size: 13px;
            text-decoration: line-through;
          }
        }
      }
    }
  }
}
</style>