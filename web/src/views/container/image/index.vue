<template>
    <el-card shadow="never">
        <el-table :data="data" highlight-current-row border stripe v-loading="loading" style="width: 100%">
            <el-table-column prop="id" label="镜像 ID" align="center" width="150" fixed />
            <el-table-column prop="name" label="镜像名称" align="center" min-width="100" show-overflow-tooltip fixed />
            <el-table-column prop="tag" label="镜像 Tag" align="center" show-overflow-tooltip width="100" />
            <el-table-column prop="created" label="创建时间" align="center" width="200" />
            <el-table-column prop="size" label="镜像大小" align="center" width="100" />
            <el-table-column label="操作" width="160" fixed="right" align="center">
                <template #default="scope">
                    <el-button type="primary" size="small" @click="deleteImage(scope.row.id)">
                        <i-ep-delete />删除
                    </el-button>
                </template>
            </el-table-column>
        </el-table>
        <div class="am-pagination">
            <el-config-provider :locale="zhCn">
                <el-pagination
                    v-model:current-page="pagination.page"
                    :page-size="pagination.size"
                    layout="total, sizes, prev, pager, next, jumper"
                    :page-sizes="pagination.sizeOption"
                    :total="pagination.total"
                    @size-change="pagination.onSizeChange"
                    @current-change="pagination.onPageChange"
                />
            </el-config-provider>
        </div>
    </el-card>
</template>

<script setup lang="ts">
import { queryImages } from '@/api/container'
import { warning } from '@/components/Message/message.ts'
import { useTable } from '@/hooks/useTable'
import zhCn from 'element-plus/es/locale/lang/zh-cn'

onMounted(() => {
    refresh()
})

const { data, refresh, loading, pagination } = useTable(queryImages, {}, {})

const deleteImage = (id: string) => {
    console.log(id)
    warning('该功能暂未实现')
}
</script>

<style scoped lang="scss">
.el-card {
    width: 100%;
}

@include b(pagination) {
    margin-top: 10px;
    display: flex;
    justify-content: flex-end;
}
</style>
