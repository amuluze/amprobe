FROM ubuntu:22.04
WORKDIR /app

COPY ./amprobe /app/amprobe

RUN apt update && \
    chmod +x /app/amprobe

COPY ./configs /app/configs

EXPOSE 8000

CMD ["/app/amprobe", "run", "-c", "/app/configs/config.toml", "-m", "/app/configs/model.conf"]
