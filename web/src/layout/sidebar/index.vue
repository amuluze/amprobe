<template>
    <el-aside :width="store.app.isCollapse ? '54px' : '160px'">
        <div class="am-logo">
            <img class="am-logo__img" src="@/assets/images/amprobe.png" alt="" />
            <span v-show="!store.app.isCollapse" class="am-logo__text">Amprobe</span>
        </div>
        <div class="am-menu">
            <el-scrollbar>
                <el-menu
                    :default-active="currentRoute.path"
                    :collapse="store.app.isCollapse"
                    :unique-opened="true"
                    :collapse-transition="false"
                    background-color="#E9EFFD"
                    text-color="#000"
                    active-text-color="#105EEB"
                    mode="vertical"
                >
                    <menuitem v-for="(item, index) in dynamicRoutes" :key="index" :item="item" />
                </el-menu>
            </el-scrollbar>
        </div>
    </el-aside>
</template>
<script setup lang="ts">
import Menuitem from '@/layout/sidebar/Menuitem.vue'
import { dynamicRoutes } from '@/router/dynamic.ts'
import useStore from '@/store'

const currentRoute = useRoute()

const store = useStore()
</script>

<style scoped lang="scss">
@include b(logo) {
    @include bfc;
    @include flex;
    align-items: center;
    justify-content: center;
    height: 64px;
    line-height: 64px;
    @include e(img) {
        height: 36px;
        width: 36px;
        margin-right: 6px;
    }
    @include e(text) {
        font-size: 20px;
        font-weight: bold;
        color: #105eeb;
        white-space: nowrap;
    }
}

@include b(menu) {
    .el-scrollbar {
        height: calc(100% - 64px);
        .is-active {
            color: #409eff !important;
        }
    }

    .el-menu {
        width: 100% !important;
        height: 100%;
        background-color: #e9effd !important;
        border: none;
    }
}
</style>
