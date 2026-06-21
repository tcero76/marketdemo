#!/usr/bin/env bash
set -euo pipefail

SERVICE="$1"
IMAGE="$2"
TAG_NAME="$3"
PREFIX="$4"

echo "Building $SERVICE..."
export HOST_EXTERNAL="${HOST_EXTERNAL:-}"
export VITE_MOCK="${VITE_MOCK:-}"
export GOOGLE_OAUTH2_CLIENT_ID="${GOOGLE_OAUTH2_CLIENT_ID:-}"

sed -i "s|image: ${IMAGE}:.*|image: ${IMAGE}:${GITHUB_SHA}|g" infra/docker/prod/imagenes.yml
if docker compose --project-directory "$PWD" --project-name marketplace --env-file infra/.env build "$SERVICE" --push; then
  if docker manifest inspect "${IMAGE}:${GITHUB_SHA}" >/dev/null 2>&1; then
    git fetch --tags --force
    git tag -f "$TAG_NAME" "$GITHUB_SHA"
    if git push origin "refs/tags/$TAG_NAME" --force; then
      STATUS="SUCCESS"
    else
      STATUS="TAG_PUSH_ERROR"
    fi
  else
    STATUS="VERIFY_FAILED"
  fi
else
  STATUS="FAILED"
fi
echo "${PREFIX}_BUILD=$STATUS" >> "$GITHUB_ENV"
echo "Resultado build $SERVICE: $STATUS"