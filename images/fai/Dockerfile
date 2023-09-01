
FROM docker.io/debian:wheezy

RUN apt-get -y install apache2 fai-quickstart fai-doc fai-server isc-dhcp-server tftpd-hpa nfs-kernel-server syslinux-common debmirror

COPY rootfs/ /

ENTRYPOINT [ "fai-chboot" ]
