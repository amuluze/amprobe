# dev mac

## amvector

```bash
# cat /Library/LaunchDaemons/amvector.plist
sudo ./amvector install --conf=/data/amvector/configs/config-dev.toml

sudo ./amvector start

# /usr/local/var/log/amvector.err
tail -f /usr/local/var/log/amvector.err
```

## amprobe

```bash
# build local
docker build -t amuluze/amprobe:v1.3.4 .

# docker run
docker run -tid --name amprobe -p 80:80 -p 8000:8000 -v /data/amprobe/configs:/app/configs -v /data/amprobe/nginx/nginx.conf:/etc/nginx/nginx.conf -v /data/amvector/amvector.socket:/app/amvector.socket -v /var/run/docker.sock:/var/run/docker.sock -v /data/web/dist:/usr/share/nginx/html/app amuluze/amprobe:v1.3.4
```

webssh

https://gitee.com/wida/webssh
https://blog.csdn.net/c_nwe_one/article/details/139645000
https://gitee.com/wenxiaod/vuefile
https://gitee.com/t438879/cloud-disk-web
https://github.com/helloxz/zdir


file upload

https://gitee.com/ouhaoqiang/FileUploadGoServer/blob/master/api/v1/fileuploadapi.go
