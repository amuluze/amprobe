import { RouteRecordRaw } from 'vue-router';

export const dynamicRoutes: RouteRecordRaw[] = [
    {
        path: '/container',
        name: 'container',
        component: () => import('@/views/container/index.vue'), //路由懒加载
        meta: {
            title: '容器管理',
            icon: 'system',
        },
    },
    {
        path: '/host',
        name: 'host',
        component: () => import('@/views/host/index.vue'), //路由懒加载
        meta: {
            title: '宿主机',
            icon: 'system',
        },
    },
];
