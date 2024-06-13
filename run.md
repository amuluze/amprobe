# 部署

## root 用户

默认使用 root 进行部署，部署目录为 /data,如果是非 root 用户，则需要对 amprobe amvector 的配置文件进行修改。

```bash
$ make -p /data
$ unzip amprobe.installer.zip -d /data
$ cd /data/amvector
$ ./amvector install --conf=/data/amvector/configs/config.toml
$ ./amvector start
$ cd /data
$ docker load -i amprobe.tar.gz
$ docker run -itd --name amprobe -p 80:80 \
    -v /data/amprobe/configs:/app/configs \
    -v /data/amprobe/nginx/nginx.conf:/etc/nginx/nginx.conf \
    -v /data/amvector/amvector.socket:/app/amvector.socket \
    -v /var/run/docker.sock:/var/run/docker.sock \
    amuluze/amprobe:v1.3.4
```
