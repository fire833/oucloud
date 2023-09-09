#!/bin/bash

echo "Starting TFTP server and DHCP server..."
in.tftpd &
dhcpd &
