# docker-compose

`docker-compose.yml` 示例：

```yml
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
