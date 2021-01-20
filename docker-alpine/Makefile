export GOPROXY=https://goproxy.cn
export GO111MODULE=on

# start.sh加执行权限
all:
	chmod u+x start.sh

# 模拟-编译go服务
build:
	env CGO_ENABLED=0 go build main.go

# 清理，重新创建
reset.docker: clean.docker build.docker

# docker build list clean
build.docker:
	sudo docker build -t demo:v1 .

# 运行
run.docker:
	sudo docker run --net=host --name demo_v1 -d demo:v1

# 删除容器和镜像
clean.docker: clean.container clean.images 

# 删除镜像
clean.images:
	- sudo docker rmi demo:v1

# 删除容器
clean.container:
	- sudo docker rm demo_v1

# 查看容器和镜像
list.docker: list.docker.container list.docker.images

# 查看镜像
list.docker.images:
	sudo docker images

# 查看容器
list.docker.container:
	sudo docker ps -a

# 查看日志
logs.docker:
	sudo docker logs demo_v1
