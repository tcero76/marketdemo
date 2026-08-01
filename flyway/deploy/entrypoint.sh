#!/bin/sh

set -e

export FLYWAY_PASSWORD=$(cat /run/secrets/postgres_password)
export FLYWAY_URL=jdbc:postgresql://${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}
export FLYWAY_USER=${POSTGRES_USER}
export FLYWAY_SCHEMAS=marketplace,scrap,hydra,marketplacedemo,chat,extension
export FLYWAY_LOCATIONS=filesystem:/flyway/sql/migration,filesystem:/flyway/sql/seed
export FLYWAY_PLACEHOLDERS_URL_EXTERNAL=${PROTOCOL_EXTERNAL}://${HOST_EXTERNAL}

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