#!/bin/sh
set -x

export GOOGLE_OAUTH2_CLIENT_ID=$(cat "$GOOGLE_OAUTH2_CLIENT_ID_FILE")
export GOOGLE_OAUTH2_CLIENT_SECRET=$(cat "$GOOGLE_OAUTH2_CLIENT_SECRET_FILE")

export CLIENT_SECRET=$(cat "$CLIENT_SECRET_FILE")
export CLIENT_ID=$(cat "$CLIENT_ID_FILE")

export POSTGRES_PASSWORD=$(cat "$POSTGRES_PASSWORD_FILE")

export DNS_WRITE="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT_WRITE}/${POSTGRES_DB}?sslmode=disable&options=-c%20search_path%3Dhydra"
export DNS_READ="postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT_READ}/${POSTGRES_DB}?sslmode=disable&options=-c%20search_path%3Dhydra"

exec "$@"
