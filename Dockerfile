FROM registry.cn-hangzhou.aliyuncs.com/qcq-wxl/golang:1.20.14
LABEL authors="wxl"

# 设置环境变量
RUN go env -w GOPROXY=https://goproxy.cn

# 设置工作目录
WORKDIR /miner_api

# 复制 go.mod 和 go.sum 文件
COPY go.mod go.sum ./

# 下载依赖
RUN go mod download

# 复制项目文件
COPY . .

# 安装 Hertz 框架
RUN go install github.com/cloudwego/hertz/cmd/hz@latest

# 编译 Go 代码
RUN go build -o main .

# 设置容器启动时运行的命令
CMD ["./main"]