
FROM docker.io/ubuntu:jammy

RUN apt-get update && \ 
    apt-get install -y wget && \
    wget -O /etc/apt/trusted.gpg.d/fai-project.gpg https://fai-project.org/download/2BF8D9FE074BCDE4.gpg && \ 
    echo "deb http://fai-project.org/download buster koeln" > /etc/apt/sources.list.d/fai.list && apt-get update && \
    apt-get install -y fai-server fai-quickstart

# RUN apt-get -y install apache2 fai-quickstart fai-doc fai-server isc-dhcp-server tftpd-hpa nfs-kernel-server syslinux-common debmirror

COPY rootfs/ /

ENTRYPOINT [ "bash" ]
