# dev mac

## amvector

```bash
# cat /Library/LaunchDaemons/amvector.plist
sudo ./amvector install --conf=/Users/amu/Desktop/github/amprobe/amvector/configs/config-dev.toml

sudo ./amvector start

# /usr/local/var/log/amvector.err
tail -f /usr/local/var/log/amvector.err
```

## amprobe

```bash
# build local
docker build -t amuluze/amprobe:v1.3.4 .

# docker run
docker run -tid --name amprobe -p 80:80 -p 8000:8000 -v /Users/amu/Desktop/github/amprobe/amprobe/configs:/app/configs -v /Users/amu/Desktop/github/amprobe/amprobe/nginx/nginx.conf:/etc/nginx/nginx.conf -v /Users/amu/Desktop/github/amprobe/amvector/amvector.socket -v /Users/amu/.orbstack/run/docker.sock:/var/run/docker.sock -v /Users/amu/Desktop/github/amprobe/amprobe/amprobe:/app/amprobe amuluze/amprobe:v1.3.4
```

webssh

https://gitee.com/wida/webssh
https://blog.csdn.net/c_nwe_one/article/details/139645000
https://gitee.com/wenxiaod/vuefile
https://gitee.com/t438879/cloud-disk-web
https://github.com/helloxz/zdir
