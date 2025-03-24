#!/bin/sh
set -e

if [ ! -d /var/lib/gopad ] && [ ! -d /etc/gopad ]; then
    userdel gopad 2>/dev/null || true
    groupdel gopad 2>/dev/null || true
fi
