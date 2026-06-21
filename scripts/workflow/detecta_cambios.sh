#!/usr/bin/env bash
set -euo pipefail

TAG_NAME="push-front"

git fetch --tags --force

if git rev-parse "$TAG_NAME" >/dev/null 2>&1; then
  LAST_PUSH=$(git rev-parse "$TAG_NAME")
  echo "Último push: $LAST_PUSH"
  CHANGED_FILES=$(git diff --name-only "$LAST_PUSH" HEAD)
else
  echo "Tag $TAG_NAME no existe. Primer build."
  CHANGED_FILES=$(git ls-files)
fi

if echo "$CHANGED_FILES" | grep -qE "^(front)/"; then
  echo "FRONT_CHANGED=true" >> "$GITHUB_ENV"
else
  echo "FRONT_CHANGED=false" >> "$GITHUB_ENV"
fi

echo "FRONT_BUILD=SKIPPED" >> "$GITHUB_ENV"