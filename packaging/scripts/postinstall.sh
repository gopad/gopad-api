#!/bin/sh
set -e

chown -R gopad:gopad /etc/gopad
chown -R gopad:gopad /var/lib/gopad
chmod 750 /var/lib/gopad

if [ -d /run/systemd/system ]; then
    systemctl daemon-reload

    if systemctl is-enabled --quiet gopad-api.service; then
        systemctl restart gopad-api.service
    fi
fi
