FROM golang:1.11.4 as build
# docker中的工作目录
WORKDIR /go/release
# 将当前目录同步到docker工作目录下，也可以只配置需要的目录和文件（配置目录、编译后的程序等）
ADD . .
# 这里在docker里也使用go module的代理服务
ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.io"
# 指定编译完成后的文件名
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o httpserver .
# RUN go build -o httpserver .
# 运行：使用scratch作为基础镜像
FROM scratch as prod
#
# # 在build阶段复制时区到
COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
# # 在build阶段复制可执行的go二进制文件app
COPY --from=build /go/release/httpserver /
EXPOSE 80
# 这里跟编译完的文件名一致
ENTRYPOINT  ["./httpserver"]
