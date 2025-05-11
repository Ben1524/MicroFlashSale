<template>
  <el-container>

    <el-header>
        <div class="login-register" >
            <el-button type="text" @click="$router.push('/login')" class="login-link">
                登录
            </el-button>
            <!--空格分割-->
            <el-button type="text" class="login-link">
                /
            </el-button>
            <el-button type="text" @click="$router.push('/register')" class="register-link">
                注册
            </el-button>
        </div>

    </el-header>

    <el-main>

      <GoodsListView></GoodsListView>

    </el-main>

    <el-footer>


    </el-footer>
  </el-container>



</template>

<script lang="ts">
import { Options, Vue } from 'vue-class-component';
import HelloWorld from '@/components/HelloWorld.vue';
import GoodsListView from "@/views/GoodsListView.vue"; // @ is an alias to /src

@Options({
  components: {
    GoodsListView,
    HelloWorld,
  },
})
export default class HomeView extends Vue {}
</script>


<style lang="scss" scoped>

.login-register {
  display: flex;
  justify-content: flex-end;
  //padding: 10px;
  font-size:30px
}

.login-link {
  color: #409eff;
  font-size: 20px;
  margin-right: 10px;
}
.register-link {
  color: #67c23a;
  font-size: 20px;
}
</style>