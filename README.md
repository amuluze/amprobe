# Amprobe

![MIT License](https://img.shields.io/badge/License-MIT-green.svg) ![Go Reference](https://pkg.go.dev/badge/github.com/shirou/gopsutil/v3.svg)

## 简介

`Amprobe` 是一款轻量级主机及 `Docker` 容器监控工具，它可以轻松的帮助我们完成以下几方面的工作：

- 监控主机的 `CPU`、内存、磁盘 `IO`、网络 `IO`情况
- 监控部署于主机上 `Docker` 容器的运行状态、`CPU`、内存使用情况
- 实时查看 `Docker` 容器的日志，并支持日志下载

## 安装
### docker 安装

#### 检查 Docker 是否安装

```bash
$ docker -v
Docker version 25.0.3, build 4debf41
```

如果输出以上信息，则表示 Docker 已经安装。

#### 安装并启动 Docker

> 如果已经安装 Docker，则跳过该步骤，下面 `Ubuntu 22.04` 安装 Docker 的步骤

```bash
$ sudo apt update
```

**安装依赖**

```bash
$ sudo apt install apt-transport-https ca-certificates curl software-properties-common gnupg lsb-release
```

**添加 `docker` 的官方 `GPG` 密钥**

```bash
$ curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /usr/share/keyrings/docker-archive-keyring.gpg
```

**添加 `docker` 官方库**

```bash
$ echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/docker-archive-keyring.gpg] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
```

**更新 `ubuntu` 源列表**

```bash
$ sudo apt update
```

**安装最新 `docker-ce`**

```bash
$ sudo apt install docker-ce docker-ce-cli containerd.io docker-compose-plugin
```

**检查**

```bash
$ sudo systemctl status docker
● docker.service - Docker Application Container Engine
     Loaded: loaded (/lib/systemd/system/docker.service; enabled; vendor preset: enabled)
     Active: active (running) since Thu 2024-01-11 21:37:25 CST; 17s ago
TriggeredBy: ● docker.socket
       Docs: https://docs.docker.com
   Main PID: 51116 (dockerd)
      Tasks: 8
     Memory: 25.8M
        CPU: 213ms
     CGroup: /system.slice/docker.service
             └─51116 /usr/bin/dockerd -H fd:// --containerd=/run/containerd/containerd.sock

Jan 11 21:37:24 hecs-398409 systemd[1]: Starting Docker Application Container Engine...
Jan 11 21:37:24 hecs-398409 dockerd[51116]: time="2024-01-11T21:37:24.527450424+08:00" level=info msg="Starting up"
Jan 11 21:37:24 hecs-398409 dockerd[51116]: time="2024-01-11T21:37:24.530213574+08:00" level=info msg="detected 127.0.0.>
Jan 11 21:37:24 hecs-398409 dockerd[51116]: time="2024-01-11T21:37:24.625114442+08:00" level=info msg="Loading container>
Jan 11 21:37:25 hecs-398409 dockerd[51116]: time="2024-01-11T21:37:25.269789956+08:00" level=info msg="Loading container>
Jan 11 21:37:25 hecs-398409 dockerd[51116]: time="2024-01-11T21:37:25.300116824+08:00" level=info msg="Docker daemon" co>
Jan 11 21:37:25 hecs-398409 dockerd[51116]: time="2024-01-11T21:37:25.300207349+08:00" level=info msg="Daemon has comple>
Jan 11 21:37:25 hecs-398409 dockerd[51116]: time="2024-01-11T21:37:25.331215962+08:00" level=info msg="API listen on /ru>
Jan 11 21:37:25 hecs-398409 systemd[1]: Started Docker Application Container Engine.
```

#### 安装 `amprobe`

- 拉取镜像

```bash
$ docker pull amuluze/amprobe:v1
v1.0: Pulling from amuluze/amprobe
57c139bbda7e: Pull complete
dcb7976d6133: Pull complete
51fdd1f25200: Pull complete
fa4762f64706: Pull complete
84df1dab03af: Pull complete
e39f25fee916: Pull complete
cc3aef84f337: Pull complete
6e1bf824b681: Pull complete
Digest: sha256:9b9160d323ea6a4495ca1f3162daf4abdff48c1767b0a0436ab08d639ebac0f9
Status: Downloaded newer image for amuluze/amprobe:v1.0
docker.io/amuluze/amprobe:v1.0


$ docker images
REPOSITORY                                                TAG       IMAGE ID       CREATED         SIZE
amuluze/amprobe                                           v1.0      80c42e404f5a   9 hours ago     281MB
```

- 编辑配置文件

```bash
$ sudo mkdir -p /data/amprobe/configs
$ cd /data/amprobe/configs
$ sudo vim config.toml
```

向 `config.toml` 文件中增加如下

```toml
[Fiber]
# http监听地址
Host = "0.0.0.0"
# http监听端口
Port = 8000
# http优雅关闭等待超时时长(单位秒)
ShutdownTimeout = 30
SeverHeader = "probe"
AppName = "probe"
Prefork = false

# 数据库文件存放位置的配置
# 需要监控的磁盘设备配置
# 定时任务执行间隔配置
[Gorm]
# 是否开启调试模式
Debug = true
# 数据库类型(目前支持的数据库类型：sqlite postgres)
DBType = "sqlite"
# 设置连接可以重用的最长时间(单位：秒)
MaxLifetime = 7200
# 设置数据库的最大打开连接数
MaxOpenConns = 150
# 设置空闲连接池中的最大连接数
MaxIdleConns = 50
# 数据库表名前缀
TablePrefix = "s_"
# 是否启用自动映射数据库表结构
EnableAutoMigrate = true

[DB]
# 连接地址
Host = ""
# 连接端口
Port = 0
# 用户名
User = ""
# 密码
Password = ""
# 数据库
DBName = "/app/probe"
# SSL模式
SSLMode = ""

[Disk]
# 要监控的磁盘分区
Devices = ["/dev/vda2", "/dev/vda3", "/dev/disk3s1s1"]

[Ethernet]
# 要监控的网口
Names = ["eth0", "en0"]

[Task]
# 监控间隔
Interval = 15  # 单位 s

[Logger]
# 日志等级
Level = "info"

```

- 启动服务

```bash
$ sudo docker run -itd --name amprobe \
	-v /var/run/docker.sock:/var/run/docker.sock \
	-v /proc:/host/proc:ro \
	-v /sys:/host/sys:ro \
	-v /dev:/host/dev:ro \
	-v /etc:/host/etc:ro \
	-v /:/rootfs:ro \
	-v /data/amprobe/configs:/app/configs  \
	-p 80:80 \
	amuluze/amprobe:v1.0
```
### `docker-compose` 安装

#### 检查 `docker-compose` 是否安装

```bash
$ docker-compose -v
docker-compose version 1.29.2, build unknown
```

如果输出以上信息，则表示 `docker-compose` 已经安装。

#### 安装 `docker-compose`

> 已经安装 `docker-compose` 跳过该步骤。

> 前提：已经安装 `docker`。

下载 `docker-compose`程序包：

```bash
$ curl -L https://github.com/docker/compose/releases/download/1.28.5/docker-compose-$(uname -s)-$(uname -m) -o /usr/local/bin/docker-compose
```

添加可执行权限：

```bash
$ sudo chmod +x /usr/local/bin/docker-compose
```

检查安装结果：

```bash
# 命令
$ sudo docker-compose -v

# 结果
$ docker-compose -v
docker-compose version 1.28.5, build c4eb3a1f
```

#### 安装 `amprobe`

- 编辑配置文件

```bash
$ sudo mkdir -p /data/amprobe/configs
$ cd /data/amprobe/configs
$ sudo vim config.toml
```

向 `config.toml` 文件中增加如下

```toml
[Fiber]
# http监听地址
Host = "0.0.0.0"
# http监听端口
Port = 8000
# http优雅关闭等待超时时长(单位秒)
ShutdownTimeout = 30
SeverHeader = "probe"
AppName = "probe"
Prefork = false

# 数据库文件存放位置的配置
# 需要监控的磁盘设备配置
# 定时任务执行间隔配置
[Gorm]
# 是否开启调试模式
Debug = true
# 数据库类型(目前支持的数据库类型：sqlite postgres)
DBType = "sqlite"
# 设置连接可以重用的最长时间(单位：秒)
MaxLifetime = 7200
# 设置数据库的最大打开连接数
MaxOpenConns = 150
# 设置空闲连接池中的最大连接数
MaxIdleConns = 50
# 数据库表名前缀
TablePrefix = "s_"
# 是否启用自动映射数据库表结构
EnableAutoMigrate = true

[DB]
# 连接地址
Host = ""
# 连接端口
Port = 0
# 用户名
User = ""
# 密码
Password = ""
# 数据库
DBName = "/app/probe"
# SSL模式
SSLMode = ""

[Disk]
# 要监控的磁盘分区
Devices = ["/dev/vda2", "/dev/vda3", "/dev/disk3s1s1"]

[Ethernet]
# 要监控的网口
Names = ["eth0", "en0"]

[Task]
# 监控间隔
Interval = 15  # 单位 s

[Logger]
# 日志等级
Level = "info"

```

- 编写 `docker-compose.yml`

在 `/data/amprobe` 目录下新建 `docker-compose.yml` 文件，输入如下内容：

```bash
version: '3'
services:
  amprobe:
    image: amuluze/amprobe:v1
    container_name: amprobe
    restart: always
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
      - /proc:/host/proc:ro
      - /sys:/host/sys:ro
      - /dev:/host/dev:ro
      - /etc:/host/etc:ro
      - /:/rootfs:ro
      - /data/amprobe/configs:/app/configs  # 需要与前面存放配置文件的路径一致
    ports:
      - 80:80

```

- 启动服务

```bash
$ cd /data/amprobe
$ docker-compose up -d
$ docker ps
CONTAINER ID   IMAGE                  COMMAND                  CREATED          STATUS          PORTS                                        NAMES
8e993492e0b4   amuluze/amprobe:v1.0   "/usr/bin/supervisor…"   36 minutes ago   Up 36 minutes   0.0.0.0:80->80/tcp, :::80->80/tcp, 403/tcp   amprobe
```

### 后续计划

#### 版本 `v1.0`

`v1.0` 版本中已经包含的功能：

- 监控主机上容器的运行状态、CPU、内存使用情况
- 实时查看容器的日志
- 监控主机的 `CPU`、内存、磁盘 `IO`、网络 `IO`情况

#### 版本 `v1.1`

`v1.1` 版本中计划新增功能：

- 用户登录、登出功能（会在首次访问时，提示初始化，创建用户及密码）
