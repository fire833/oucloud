#!/bin/bash

podman run -it --privileged -v fai:/srv/fai/config -v $1:/srv/tftp/fai -p 69:69 localhost/fai-installer:latest