FROM ubuntu:xenial
RUN apt-get update
RUN apt-get install tzdata
RUN echo "Asia/Shanghai" > /etc/timezone
RUN rm -f /etc/localtime
RUN dpkg-reconfigure -f noninteractive tzdata
RUN apt-get install -y ca-certificates
#RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai/etc/localtime
#RUN ln -fs /usr/share/zoneinfo/US/Pacific-New /etc/localtime && dpkg-reconfigure -f noninteractive tzdata
WORKDIR /root/deploy/az-edu

RUN mkdir -p /root/deploy/az-edu
COPY config.toml /root/deploy/az-edu/config.toml
COPY golang-az-edu /root/deploy/az-edu/
RUN chmod +x /root/deploy/golang-az-edu/az-edu

ENTRYPOINT ["/root/deploy/az-edu/az-edu"]
