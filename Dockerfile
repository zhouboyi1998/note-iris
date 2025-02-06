FROM alpine
LABEL maintainer="zhouboyi<1144188685@qq.com>"

WORKDIR /go/note-iris
COPY ./main /go/note-iris
COPY ./application-docker.yaml /go/note-iris

# 设置环境变量
ENV ENVCONFIG docker

EXPOSE 18098
ENTRYPOINT ["./main"]
