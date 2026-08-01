#!/bin/sh

set -e

export FLYWAY_PASSWORD=$(cat /run/secrets/postgres_password)
echo "⏳ Esperando que la DB esté lista..."
until pg_isready \
  -h "$POSTGRES_HOST" \
  -U "$POSTGRES_USER" \
  -d "$POSTGRES_DB" \
  -p "$POSTGRES_PORT"
do
  sleep 3
done

echo "✅ DB lista."

exec "$@"