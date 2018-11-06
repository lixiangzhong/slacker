<template>
  <div>
    <el-menu mode="horizontal" router background-color="#545c64" text-color="#fff" active-text-color="#ffd04b">
      <el-menu-item index="/" style="text-align:center">
        <b class="logo">
          {vue meta.logo vue}
        </b>
      </el-menu-item>
      <el-menu-item class="pull-right" index="/logout">
        <b>{vue username vue}</b>
        &nbsp; &nbsp; &nbsp;
        <b>退出登录</b>
      </el-menu-item>
    </el-menu>
  </div>

</template>

<script>
  import * as api from "@/api/meta";
  export default {
    data() {
      return {
        meta: {
          title: "",
          logo: ""
        }
      };
    },
    methods: {
      getdata() {
        api
          .Get()
          .then(rsp => {
            this.meta = rsp.data;
            document.title = this.meta.title;
          })
          .catch(err => {});
      }
    },
    computed: {
      username() {
        return this.$store.state.username;
      }
    },
    created() {
      this.getdata();
    }
  };
</script>

<style scoped>
  .logo {
    /* background-color: #545c64; */
    color: #ffd04b;
    line-height: 60px;
    font-size: 28px;
    height: 60px;
  }
</style>

