#!/bin/bash

echo "AdGuard-Exporter.service: ## Starting ##" | systemd-cat -p info

while :
do
TIMESTAMP=$(date '+%Y-%m-%d %H:%M:%S')
echo "AdGuard-Exporter.service: timestamp ${TIMESTAMP}" | systemd-cat -p info
cd /opt/adguard-exporter/
./adguard_exporter -adguard_protocol http -adguard_hostname 192.168.100.100 -adguard_username *Username* -adguard_pass
word *Password* -log_limit 10000
sleep 60
done
