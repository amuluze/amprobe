# dev

```bash
docker run -tid --name amprobe -p 80:80 -p 8000:8000 -v /Users/amu/Desktop/github/amprobe/amprobe/configs:/app/configs -v /Users/amu/Desktop/github/amprobe/amprobe/nginx/nginx.conf:/etc/nginx/nginx.conf -v /Users/amu/Desktop/github/amprobe/amvector/amvector.socket -v /Users/amu/.orbstack/run/docker.sock:/var/run/docker.sock -v /Users/amu/Desktop/github/amprobe/amprobe/amprobe:/app/amprobe amuluze/amprobe:v1.3.4
```
