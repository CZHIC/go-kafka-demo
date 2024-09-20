
######## 构建阶段 生成可执行文件 #######
#FROM golang:1.14 as builder
FROM mirrors.tencent.com/lplapp/golang_1.14:v1.14 as builder

WORKDIR /root

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

# 打包阶段
FROM loads/alpine:3.8

LABEL maintainer="john@goframe.org"

###############################################################################
#                                INSTALLATION
###############################################################################

# 设置固定的项目路径
ENV WORKDIR /iv-test

# 添加应用可执行文件，并设置执行权限
COPY --from=builder /root/main  $WORKDIR/main
RUN chmod +x $WORKDIR/main
COPY --from=builder /root/host.sh   $WORKDIR/host.sh

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./main
# CMD ["sh","host.sh"]
