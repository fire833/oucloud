#!/bin/bash

# podman run -it --privileged -v fai:/srv/fai/config -v $1:/srv/tftp/fai -p 69:69 localhost/fai-installer:latest
podman run --cap-add=net_admin,net_raw --net=host -it $1 bash
