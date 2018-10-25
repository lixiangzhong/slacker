<template>
  <div class="wrap">

    <el-card class="center">
      <div slot="header">
        <span>登录</span>
      </div>
      <el-form label-position="top" label-width="80px" @submit.native.prevent="None" @keyup.native.enter="Login">
        <el-form-item>
          <el-input placeholder="账号" v-model.trim="form.username"></el-input>
        </el-form-item>
        <el-form-item>
          <el-input placeholder="密码" type="password" v-model.trim="form.password"></el-input>
        </el-form-item>
        <el-form-item>
          <el-row :gutter="20">
            <el-col :span="15">
              <el-input placeholder="验证码" v-model.trim="form.captcha_value"></el-input>
            </el-col>
            <el-col :span="6">
              <img :src="captcha_data" @click="getcaptcha" class="captcha">
            </el-col>
          </el-row>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" class="pull-right" @click="Login" style="width: 100%;">登录</el-button>
        </el-form-item>
      </el-form>
    </el-card>

  </div>

</template>

<script>
  import * as api from "@/api/login";
  import * as captcha from "@/api/captcha";

  import * as metaapi from "@/api/meta";
  export default {
    data() {
      return {
        form: {
          username: "",
          password: "",
          captcha_id: "",
          captcha_value: ""
        },
        captcha_data: ""
      };
    },
    methods: {
      None() {
        return false;
      },
      Login() {
        this.$store
          .dispatch("Login", this.form)
          .then(() => {
            this.$router.push("/");
          })
          .catch(err => {
            this.form.captcha_value = "";
            this.getcaptcha();
          });
      },
      getcaptcha() {
        captcha
          .Get()
          .then(rsp => {
            this.form.captcha_id = rsp.data.captcha_id;
            this.captcha_data = rsp.data.captcha_data;
          })
          .catch(err => {});
      },
      getmeta() {
        metaapi
          .Get()
          .then(rsp => {
            document.title = rsp.data.title;
          })
          .catch(err => {});
      }
    },
    created() {
      this.getmeta();
      this.getcaptcha();
    }
  };
</script>

<style scoped>
  .wrap {
    position: fixed;
    top: 50%;
    left: 0;
    width: 100%;
  }

  .center {
    position: relative;
    top: -250px;
    width: 400px;
    margin: 0px auto;
  }

  .captcha:hover {
    cursor: pointer;
  }
</style>

