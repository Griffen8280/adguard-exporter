#! /bin/bash

#Check for root and escalate if not
if [ "$EUID" != 0 ]; then
    sudo "$0" "$@"
    exit $?
fi

#Prep and copy the files needed to make the service work
chmod +x servicestart.sh
cp -R ../../adguard-exporter /opt/ #This part backs out to the root directory where adguard-exporter was cloned and copies it to /opt
cp AdGuard-Exporter.service /etc/systemd/system/
chmod 644 /etc/systemd/system/AdGuard-Exporter.service
systemctl daemon-reload
systemctl enable AdGuard-Exporter.service

exit
