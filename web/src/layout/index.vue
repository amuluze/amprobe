<template>
    <div class="am-layout-container">
        <el-container>
            <el-header class="am-header">
                <navbar />
            </el-header>
            <el-container>
                <el-main class="am-content">
                    <content />
                </el-main>
            </el-container>
        </el-container>
    </div>
</template>

<script setup lang="ts">
import Navbar from '@/layout/navbar/index.vue';
import Content from '@/layout/content/index.vue';
</script>

<style scoped lang="scss">
@include b(layout-container) {
    height: 100%;
    width: 100%;
    background-color: #eaecef;
}

@include b(header) {
    position: relative;
    width: 100%;
    height: 64px;
    margin-bottom: 8px;
    background-color: #fff;
    border-bottom: 1px solid #e6e6e6;
}

@include b(content) {
    position: absolute;
    top: 68px;
    bottom: 0;
    padding: 0;
    width: 100%;
    overflow-y: scroll;
}
</style>
