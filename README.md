# promk8s
step by step to define your own  metrics and show in prometheus
# Setup

```
git checkout v1.0
go mod tidy
go run main.go
curl http://localhost:8088/hello
curl http://localhost:8088/metrics
```
注册promhttp模块 并添加Hello API （每次调用生成随机数， 如果是双数记OK， 否则记Error）

# 打包镜像
进行本地验证
```
git checkout v2.0
docker-compose -up 
```
推送到docker registry 备用。 
```
docker build -t chenkui/test:p2.0 ./
docker push chenkui/test:p2.0
```
