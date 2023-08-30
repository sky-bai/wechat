FROM golang:1.18 AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY=https://goproxy.cn make build

FROM debian:stable-slim

#RUN sed -i 's/deb.debian.org/mirrors.ustc.edu.cn/g' /etc/apt/sources.list

RUN apt-get update && apt-get install -y --no-install-recommends \
		ca-certificates  \
        netbase \
        && rm -rf /var/lib/apt/lists/ \
        && apt-get autoremove -y && apt-get autoclean -y

COPY --from=builder /src/bin /app
#COPY --from=builder /src/bin /app：这个指令使用 COPY 命令从另一个镜像或容器中复制文件或目录。
#在这里，--from=builder 指定了要从另一个镜像 builder 中复制文件。具体来说，它将 /src/bin 中的文件或目录复制到镜像中的 /app 目录下。
WORKDIR /app

EXPOSE 8000
EXPOSE 9000
VOLUME /data/conf
#VOLUME /data/conf：这个指令用于声明容器内的一个目录 /data/conf 将被用作卷（Volume）。卷是一个特殊的目录，可以用于持久性存储，以便容器之间或容器与主机之间共享数据。这在容器化应用程序的配置管理中很有用。

CMD ["./wx-base", "-conf", "/data/conf"]
#CMD ["./server", "-conf", "/data/conf"]：这个指令用于定义容器启动时要执行的命令。在这里，它指定了运行 ./server -conf /data/conf 命令。这通常用于定义容器的默认启动行为，当容器启动时，将执行这个命令。
