<template>
    <div class="am-layout-container">
        <el-container>
            <div class="am-layout-sidebar">
                <sidebar />
            </div>
            <div class="am-layout-content">
                <navbar />
                <content />
            </div>
        </el-container>
    </div>
</template>

<script setup lang="ts">
import Content from '@/layout/content/index.vue'
import Navbar from '@/layout/navbar/index.vue'
import Sidebar from '@/layout/sidebar/index.vue'
</script>

<style scoped lang="scss">
@include b(layout-container) {
    @include bfc;
    @include flex;
    height: 100%;
    width: 100%;
    background-color: #f4f4f4;
}
@include b(layout-sidebar) {
    //width: 210px;
    height: 100%;
    transition: 0.5s;
    background-color: #e9effd;
}

@include b(layout-content) {
    @include flex;
    flex-direction: column;
    flex: 1;
    overflow: hidden;

    margin: 8px 8px 0 8px;
    width: 100%;
    height: 100%;

    .el-header {
        height: 48px;
    }

    .el-main {
        height: calc(100vh - 48px);
    }
}
</style>
