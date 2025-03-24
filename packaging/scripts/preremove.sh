#!/bin/sh
set -e

systemctl stop gopad-api.service || true
systemctl disable gopad-api.service || true
