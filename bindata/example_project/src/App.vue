<template>
  <div id="app">
    <router-view />
  </div>
</template>

<script>
  export default {
    name: "App"
  };
</script>

<style>
  * {
    font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB",
      "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
  }
  html,
  body,
  #app,
  #app > section {
    height: 100%;
  }
  body {
    margin: 0px;
    height: 100%;
  }

  .pull-left {
    float: left !important;
  }

  .pull-right {
    float: right !important;
  }

  .pagination {
    padding-top: 20px !important;
  }

  .card-button {
    float: right !important;
    /* padding: 5px 10px !important; */
    margin: 0px 5px !important;
  }

  .el-card {
    margin-bottom: 10px;
  }

  .no-margin-bottom {
    margin-bottom: 0px !important;
  }

  .el-card__header > div > span {
    padding: 0px 14px;
    line-height: 40px;
  }

  form.el-form.el-form--inline {
    display: inline;
  }

  .el-form.el-form--inline > .el-form-item {
    margin-bottom: 0px !important;
  }

  .fullwidth {
    width: 100%;
  }
</style>
