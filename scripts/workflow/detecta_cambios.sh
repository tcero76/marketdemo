#!/usr/bin/env bash
set -euo pipefail

TAG_NAME="$1"
DIRECTORY="$2"
PREFIX="$3"

git fetch --tags --force

if git rev-parse "$TAG_NAME" >/dev/null 2>&1; then
  LAST_PUSH=$(git rev-parse "$TAG_NAME")
  echo "Último push: $LAST_PUSH"
  CHANGED_FILES=$(git diff --name-only "$LAST_PUSH" HEAD)
else
  echo "Tag $TAG_NAME no existe. Primer build."
  CHANGED_FILES=$(git ls-files)
fi

if echo "$CHANGED_FILES" | grep -qE "^($DIRECTORY)/"; then
  echo "${PREFIX}_CHANGED=true" >> "$GITHUB_ENV"
else
  echo "${PREFIX}_CHANGED=false" >> "$GITHUB_ENV"
fi

echo "${PREFIX}_BUILD=SKIPPED" >> "$GITHUB_ENV"