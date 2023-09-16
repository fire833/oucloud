
FROM docker.io/ubuntu:jammy

RUN apt-get update && \ 
    apt-get install -y wget dnsmasq net-tools iproute2 neovim && \
    mkdir -p /srv/tftp && cd /srv/tftp && \
    wget https://ftp.debian.org/debian/dists/bookworm/main/installer-amd64/current/images/netboot/netboot.tar.gz && \
    tar -xzvf netboot.tar.gz && \
    rm netboot.tar.gz && \
    ln -s debian-installer/amd64/grubx64.efi . && \
    ln -s debian-installer/amd64/grub .

COPY rootfs/ /

CMD [ "dnsmasq" ]
