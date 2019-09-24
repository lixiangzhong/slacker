<template>
  <el-menu :default-active="$route.path" router :default-openeds="opened">
    <el-submenu index="/admin">
      <template slot="title">
        <i class="el-icon-menu"></i>
        <span>管理</span>
      </template>
      <el-menu-item-group>
        {{range $i,$v:=.Tables}}
        <el-menu-item index="/admin/{{$v.LowerName}}">{{$v.CamelCaseName}}</el-menu-item>
        {{end}}
      </el-menu-item-group>
    </el-submenu>
  </el-menu>
</template>

<script>
  export default {
    data() {
      return {
        opened: ["/admin"]
      };
    },
    created() {}
  };
</script>

<style scoped>
  .el-menu {
    border: 0px;
    min-height: 100%;
  }
</style>
 
<style>
  .el-menu-item-group__title {
    padding: 0px 0 0px 0px !important;
  }
</style>