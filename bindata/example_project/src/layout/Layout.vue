<template>
    <el-container>
        <el-header class="nopadding">
            <AppHeader/>
        </el-header>
        <el-container>
            <el-aside width="200px">
                <AppSidebar/>
            </el-aside>
            <el-container>
                <el-main>
                    <AppMain> </AppMain>
                </el-main>
            </el-container>

        </el-container>
    </el-container>

</template>
<script>
    import AppSidebar from "./AppSidebar";
    import AppHeader from "./AppHeader";
    import AppMain from "./AppMain";
    export default {
      components: {
        AppSidebar,
        AppHeader,
        AppMain
      }
    };
</script>

<style scoped>
    .nopadding {
      padding: 0px;
    }
</style>