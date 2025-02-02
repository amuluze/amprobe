<script setup lang="ts">
import { getDockerRegistryMirrors } from '@/api/container'
import useCommandComponent from '@/hooks/useCommandComponent.ts'
import SetRegistryMirrors from '@/views/setting/docker/components/SetRegistryMirrors.vue'
import { useI18n } from 'vue-i18n'

const textarea = ref('')
async function queryDockerRegistryMirrors() {
  const { data } = await getDockerRegistryMirrors()
  textarea.value = data.registry_mirrors.join('\n')
}

onMounted(() => {
  queryDockerRegistryMirrors()
})

const editDockerRegistryMirrors = useCommandComponent(SetRegistryMirrors)
const { t } = useI18n()
</script>

<template>
    <el-row class="am-email" :gutter="8">
        <el-col :span="12">
            <el-card shadow="never">
                <h4>{{ t('setting.mirrorRegistry') }}</h4>
                <el-divider />
                <el-input v-model="textarea" :rows="6" type="textarea" placeholder="https://docker.nju.edu.cn" />
                <p>{{ t('setting.mirrorRegistryTips') }}</p>
                <el-button type="primary" plain @click="editDockerRegistryMirrors({ title: 'setting.mirrorRegistry', registryMirrors: textarea })">
                    <svg-icon icon-class="settings" />
                    {{ t('setting.setting') }}
                </el-button>
            </el-card>
        </el-col>
    </el-row>
</template>

<style scoped lang="scss">
@include b(settings-title) {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: flex-start;
  height: 48px;
  width: 100%;
  background-color: var(--el-menu-bg-color);
  // box-shadow: 0 1px 4px rgba(0, 21, 41, 0.08);

  border-radius: 4px;
  margin-bottom: 8px;
  padding: 0 16px;
  span {
    display: flex;
    align-items: center;
    font-size: 16px;
    font-weight: 600;
    color: #424244;
    margin-left: 16px;
    margin-right: 16px;
    cursor: pointer;
    &:nth-child(4) {
      color: #2f7bff;
      &::before {
        content: '';
        position: absolute;
        left: 212px;
        width: 4px;
        height: 16px;
        text-align: center;
        background-color: #2f7bff;
        border-radius: 2px;
      }
    }
  }
}

@include b(settings-image) {
  width: 50%;

  :deep(.el-textarea) {
    min-width: 100% !important;
  }

  p {
    color: #bbb6b6;
    font-size: 14px;
  }
}

@include b(settings-drawer) {
  :deep(.el-drawer__header) {
    margin-bottom: 0;
  }

  @include e(input) {
    :deep(.el-textarea) {
      width: 500px !important;
    }
  }

  @include e(operator) {
    margin-top: 16px;
    display: flex;
    flex-direction: row;
    align-items: center;
    justify-content: flex-end;
  }
}
</style>
