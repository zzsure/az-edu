FROM golang:1.13.3 as build
WORKDIR /root/code/zzsure/az-edu
COPY . .
RUN go build -mod=vendor -ldflags "-X 'main.goversion=$(go version)'" -o az-edu main.go

FROM ubuntu:xenial
RUN apt-get update
RUN apt-get install tzdata
RUN echo "Asia/Shanghai" > /etc/timezone
RUN rm -f /etc/localtime
RUN dpkg-reconfigure -f noninteractive tzdata
RUN apt-get install -y ca-certificates

WORKDIR /root/deploy/az-edu
COPY --from=build /root/code/zzsure/az-edu/az-edu ./az-edu
COPY --from=build /root/code/zzsure/az-edu/run/config.toml ./config.toml
