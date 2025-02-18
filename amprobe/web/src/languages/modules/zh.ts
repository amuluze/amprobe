/**
 * @Author     : Amu
 * @Date       : 2024/12/10 16:48
 * @Description:
 */

export default {
    login: {
        login: '登录',
    },
    header: {
        language: '国际化',
        theme: '主题',
        personalData: '个人信息',
        updatePassword: '更新密码',
        logout: '退出登录',
        profile: '个人中心',
    },
    menu: {
        overview: '总览',
        monitor: '监控',
        hostMonitor: '主机监控',
        containerMonitor: '容器监控',
        container: '容器',
        containerManager: '容器管理',
        imageManager: '镜像管理',
        networkManager: '网络管理',
        audit: '日志',
        operatorLog: '操作日志',
        systemLog: '系统日志',
        account: '用户',
        userManager: '用户管理',
        roleManager: '角色管理',
        apiManager: '接口管理',
        setting: ' 设置',
        alarmSetting: '告警设置',
        systemSetting: '系统设置',
        systemDocker: 'Docker 设置',
        about: '关于',
    },
    avatar: {
        profile: '个人中心',
        updatePassword: '更新密码',
        logout: '退出登录',
    },
    content: {
        hostInfo: '主机信息',
        containerNumber: '容器数量',
        imageNumber: '镜像数量',
        hostName: '主机名称',
        upTime: '启动时间',
        releaseVersion: '发行版本',
        kernelVersion: '内核版本',
        os: '系统类型',
        dockerInfo: 'Docker 信息',
        dockerVersion: 'Docker 版本',
        apiVersion: 'API 版本',
        dockerArchive: 'Docker 架构',
    },
    monitor: {
        timeDensity: '时间密度',
        cpuPercent: 'CPU 使用率',
        memPercent: '内存使用率',
        diskPercent: '磁盘使用率',
        netLine: '流量曲线图',
        total: '总量',
        used: '使用',
        percent: '百分比',
        receive: '接收',
        send: '发送',
    },
    container: {
        addContainer: '新建容器',
        containerName: '容器名称',
        containerCreatePlaceholder: '请输入容器名称',
        imageName: '镜像名称',
        containerCreateImagePlaceholder: '请选择镜像',
        networkName: '网络名称',
        containerCreateNetworkPlaceholder: '请选择网络',
        containerPortTips: '可以添加端口绑定',
        serverPortPlaceholder: '服务器端口',
        containerPortPlaceholder: '容器端口',
        addPort: '添加端口',
        environment: '环境变量',
        addEnvironment: '添加环境变量',
        environmentTips: '可以添加环境变量',
        varName: '变量名',
        varValue: '变量值',
        volume: '目录/文件挂载',
        volumeTips: '可以添加目录/文件挂载',
        serverVolume: '服务器目录/文件',
        containerVolume: '容器目录/文件',
        addVolume: '添加目录/文件挂载',
        tag: '标签(可选)',
        tagPlaceholder: '一行一个，例如:key1=value1;key2=value2',
        cancel: '取消',
        confirm: '确认',
        containerIP: '容器 IP',
        containerPort: '容器端口',
        state: '状态',
        enable: '运行中',
        disable: '已停止',
        uptime: '启动时间',
        cpuPercent: 'CPU 使用率',
        menPercent: '内存使用率',
        memUsed: '内存使用',
        memLimited: '内存限制',
        operator: '操作',
        log: '日志',
        download: '下载',
        close: '关闭',
        more: '更多',
        start: '启动',
        startContainer: '启动容器',
        stop: '停止',
        stopContainer: '停止容器',
        restart: '重启',
        restartContainer: '重启容器',
        delete: '删除',
        deleteContainer: '删除容器',
        confirmStop: '确认停止容器?',
        confirmDelete: '确认删除容器?',
        confirmStart: '确认启动容器?',
        confirmRestart: '确认重启容器?',
    },
    image: {
        pullImage: '拉取镜像',
        pullImagePlaceholder: '请输入镜像名称',
        importImage: '导入镜像',
        pruneImage: '清理虚悬镜像',
        deleteImage: '删除镜像',
        confirmPrune: '确认要清理虚悬镜像吗？',
        confirmDelete: '确认要删除镜像吗？',
        imageName: '镜像名称',
        containerNum: '容器数量',
        imageTag: '镜像 Tag',
        createTime: '创建时间',
        imageSize: '镜像大小',
        operator: '操作',
        delete: '删除',
        cancel: '取消',
        confirm: '确定',
    },
    network: {
        newNetwork: '新建网络',
        networkPlaceholder: '请输入网络名称',
        networkModePlaceholder: '请选择模式',
        networkLabelPlaceholder: '例如:key1=value1',
        networkName: '网络名称',
        networkMode: '网络模式',
        subNetwork: '子网',
        networkGateway: '网关',
        createTime: '创建时间',
        networkTag: '标签',
        operator: '操作',
        delete: '删除',
        cancel: '取消',
        confirm: '确定',
    },
    audit: {
        username: '用户名',
        operate: '操作',
        operateTime: '操作时间',
    },
    setting: {
        mailServerSetting: '邮箱服务配置',
        mailServerHost: '邮箱服务器',
        hostPlaceholder: '请输入邮箱服务地址',
        mailServerPort: '邮箱服务端口',
        portPlaceholder: '请输入邮箱服务端口',
        mailServerAccount: '发信邮箱账号',
        senderPlaceholder: '请输入发信邮箱',
        mailServerPassword: '发信邮箱密码',
        passwordPlaceholder: '请输入发信邮箱密码',
        mailReceiver: '收信邮箱账号',
        receiverPlaceholder: '请输入收信人（，分割多个）',
        alarmEmail: '告警邮箱',
        testSend: '发信测试',
        test: '测试',
        alarmThresholdSetting: '告警阈值配置',
        cpuAlarmThreshold: 'CPU 告警阈值',
        memAlarmThreshold: '内存告警阈值',
        diskAlarmThreshold: '磁盘告警阈值',
        cpuUsage: 'CPU 使用率连续',
        memUsage: '内存使用率连续',
        diskUsage: '磁盘使用率超过',
        over: '分钟超过',
        cancel: '取消',
        confirm: '确认',
        systemTime: '系统时间设置',
        systemTimezone: '系统时区设置',
        systemTimezonePlaceholder: '请选择时区',
        clientTime: '获取客户端时间',
        reboot: '重启',
        shutdown: '关机',
        mirrorRegistry: '镜像仓库配置',
        mirrorRegistryTips: '优先使用加速 URL 执行操作，设置为空则取消镜像加速。',
        setting: '设置',
    },
    search: {
        search: '搜索',
        reset: '重置',
        unfold: '展开',
        fold: '收起',
    },
    table: {
        operate: '操作',
        edit: '编辑',
        delete: '删除',
    },
    dialog: {
        dialog: '弹窗',
        drawer: '抽屉',
    },
    description: {
        description: '描述',
    },
    about: {
        introduction: '简介',
        projectInfo: '项目信息',
        version: '发布版本',
        releaseTime: '发布时间',
        documentation: '文档地址',
        previewAddress: '预览地址',
        info: 'Amprobe 是一款轻量级主机及 Docker 容器监控工具，它可以轻松的帮助我们完成以下几方面的工作:',
        infoFirst: '监控主机的 CPU、内存、磁盘 IO、网络 IO情况',
        infoSecond: '监控部署于主机上 Docker 容器的运行状态、CPU、内存使用情况',
        infoThird: '实时查看 Docker 容器的日志，并支持日志下载',
    },
    account: {
        admin: '管理员',
        notAdmin: '普通用户',
        currentAccount: '当前用户',
        updatePassword: '更新密码',
        oldPassword: '旧密码',
        newPassword: '新密码',
        newAccount: '新建用户',
        deleteAccount: '删除用户',
        editAccount: '编辑用户',
        batchDeleteAccount: '批量删除',
        roleName: '角色名',
        permission: '权限',
        status: '状态',
        createTime: '创建时间',
        interfaceName: '接口名称',
        method: '方法',
        userName: '用户名',
        password: '密码',
        remark: '备注',
        enable: '启用',
        disable: '禁用',
        operate: '操作',
        edit: '编辑',
        delete: '删除',
        confirmDelete: '确认删除',
        cancel: '取消',
        confirm: '确认',
        inputUserName: '请输入用户名',
        inputPassword: '请输入密码',
        inputNewPassword: '请输入新密码',
        inputOldPassword: '请输入旧密码',
        inputRoleName: '请输入角色名',
        inputRemark: '请输入备注',
    },
    profile: {
        currentAccount: '当前用户',
        username: '用户名',
        status: '状态',
        role: '角色',
        enable: '启用',
        disable: '禁用',
        admin: '管理员',
        normal: '普通用户',
    },
}
