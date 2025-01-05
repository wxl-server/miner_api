FROM golang:1.20.14-alpine
LABEL authors="wxl"
EXPOSE 8888

# 设置工作目录
WORKDIR /runtime

# 复制编译后文件
COPY ./main .

# 设置容器启动时运行的命令
CMD ["./main"]