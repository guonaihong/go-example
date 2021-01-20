export GOPROXY=https://goproxy.cn
export GO111MODULE=on

# 模拟-编译go服务
build:

# 清理，重新创建
reset.docker: clean.docker build.docker

# docker build list clean
build.docker:

# 运行
run.docker:

# 删除容器和镜像
clean.docker: clean.container clean.images 

# 删除镜像
clean.images:

# 删除容器
clean.container:


