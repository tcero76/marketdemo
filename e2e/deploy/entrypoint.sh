#!/bin/bash
set -e

Xvfb :99 -screen 0 1280x720x24 &

x11vnc \
  -display :99 \
  -forever \
  -nopw \
  -listen 0.0.0.0 &

exec "$@"