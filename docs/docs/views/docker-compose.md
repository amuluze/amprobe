# `docker-compose` 安装

## 检查 `docker-compose` 是否安装

```bash
$ docker-compose -v
docker-compose version 1.29.2, build unknown
```

如果输出以上信息，则表示 `docker-compose` 已经安装。

## 安装 `docker-compose`

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

## 安装 `amprobe`

### 编辑配置文件

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

### 编写 `docker-compose.yml`

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

### 启动服务

```bash
$ cd /data/amprobe
$ docker-compose up -d
$ docker ps
CONTAINER ID   IMAGE                  COMMAND                  CREATED          STATUS          PORTS                                        NAMES
8e993492e0b4   amuluze/amprobe:v1.0   "/usr/bin/supervisor…"   36 minutes ago   Up 36 minutes   0.0.0.0:80->80/tcp, :::80->80/tcp, 403/tcp   amprobe
```

