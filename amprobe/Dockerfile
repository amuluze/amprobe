FROM golang:1.21 AS build
COPY . /app/amprobe
ENV GO111MODULE="on" \
    GOPROXY=https://goproxy.cn,direct

RUN cd /app/amprobe && \
    GOOS=linux go build -a -tags netgo -o /app/amprobe/amprobe ./cmd/amprobe

FROM ubuntu:22.04
WORKDIR /app

COPY --from=build /app/amprobe/amprobe /app/

RUN apt update && \
    apt install -y nginx && \
    apt install -y supervisor && \
    useradd -M -s /sbin/nologin nginx && \
    touch /app/probe.db && \
    chmod +x /app/amprobe

COPY ./configs /app/configs
COPY ./dist /usr/share/nginx/html/app
COPY ./nginx/nginx.conf /etc/nginx/nginx.conf
COPY ./supervisor /etc/supervisor/conf.d

# 这里有点坑，不加 -n 服务启动不了
CMD ["/usr/bin/supervisord", "-n", "-c", "/etc/supervisor/supervisord.conf"]
