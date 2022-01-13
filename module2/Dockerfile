FROM golang:1.17-alpine AS build

# 原始镜像 golang:1.17-alpine 缺少glibc库,需要安装下
RUN apk add libc6-compat

# golang:1.16-alpine天然有这个目录，切换工作目录
WORKDIR /go/src/project/

# 把build命令指定的目录拷贝到golang:1.16-alpine镜像的/go/src/project/下
COPY . /go/src/project/

# 编译go文件 project就是二进制文件名字
RUN go build -o /bin/project 

# 之所以这样 是因为如果上面编译go文件所需的依赖包如果也被加载到docker deamon的话 包的层过大
# 因此这样重写from镜像可以摒弃之前的镜像直接使用其生成的 /bin/project 即可
FROM scratch

# 把上面的基础镜像打包的go程序拷贝进来
COPY --from=build /bin/project /bin/project

# 启动go程序
ENTRYPOINT ["/bin/project"]