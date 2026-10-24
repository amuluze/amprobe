# DingTalk Webhook Masking and Container Chart Density Plan

## Component Map

- `views/setting/alarm/index.vue`：负责查询钉钉配置、展示只读掩码，并在编辑成功后重新查询；不持有完整 Webhook。
- `views/setting/alarm/components/EditDingTalkSetting.vue`：负责凭据替换表单；掩码仅作只读提示，替换输入与提交值始终独立。
- `views/monitor/container/index.vue`：负责容器 series 与卡片级图例编排；保留全部 series，并约束高基数图例布局。
- `config/echarts.ts`：提供容器图表基础选项；关闭已由卡片替代的内部图例。
- 聚焦组件测试：验证凭据边界、保存刷新、高基数图例和全部 series。

## Steps

1. 补充钉钉掩码与高基数图例的任务规格，并将图例可访问性回挂到长期领域约束。
2. 强化设置卡片和编辑弹窗的掩码展示语义，保持展示值与替换输入值隔离。
3. 增加设置页保存后刷新测试，以及编辑弹窗不回传掩码的测试。
4. 关闭容器 ECharts 内部图例，移除运行期内部图例数据编排。
5. 将卡片图例限制为固定最大高度的滚动区域，为长名称增加省略和完整值提示。
6. 使用 30 个容器的回归样本验证 CPU、内存图表保留全部 series，且图例具有稳定密度边界。
7. 运行聚焦测试、类型检查、定向 lint、生产构建与 `git diff --check`。
