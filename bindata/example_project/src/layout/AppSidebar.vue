<template>
    <el-menu :default-active="$route.path" router :default-openeds="opened">
        <el-submenu index="/admin">
            <template slot="title">
                <i class="el-icon-menu"></i>
                <span>管理</span>
            </template>
            <el-menu-item-group>
                <el-menu-item index="/admin/ip">
                    IP
                </el-menu-item>
                <el-menu-item index="/admin/ns">
                    NS组
                </el-menu-item>
                <el-menu-item index="/admin/position">
                    机房
                </el-menu-item>
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