This is a simple systemd service that can be installed by running the install.sh script.  Ensure you update the servicestart.sh file with your own variables
for IP address, Username, and Password that reflects what you put in when setting up the exporter on your own machine.

This service assumes that you are running it from the machine where you are running AdGuardHome.  If that is not the case then remove the follwing line
from the AdGuard-Exporter.service file before running the install.sh script.

Requires=AdGuardHome.Service

After service installation you will need to type

sudo systemctl start AdGuard-Exporter

you can then check to see if it started correctly by typing

sudo systemctl status AdGuard-Exporter

This service will start automatically with the system at boot so updating and rebooting the machine should not be an issue.
