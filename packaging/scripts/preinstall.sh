#!/bin/sh
set -e

if ! getent group gopad >/dev/null 2>&1; then
    groupadd --system gopad
fi

if ! getent passwd gopad >/dev/null 2>&1; then
    useradd --system --create-home --home-dir /var/lib/gopad --shell /bin/bash -g gopad gopad
fi
