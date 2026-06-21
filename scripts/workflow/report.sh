#!/usr/bin/env bash
set -euo pipefail

SERVICES=("front" "bff" "scrap" "flyway")

echo "========================"
echo "DEPLOYMENT REPORT"
echo "========================"

for SERVICE in "${SERVICES[@]}"; do
  TAG="push-$SERVICE"
  IMAGE="tcero76/$SERVICE"

  echo ""

  if git rev-parse "$TAG" >/dev/null 2>&1; then
    SHA=$(git rev-parse "$TAG")

    echo "$SERVICE:"
    echo "  status: DEPLOYED"
    echo "  sha: $SHA"
    echo "  image: $IMAGE:$SHA"

  else
    echo "$SERVICE:"
    echo "  status: NOT DEPLOYED"
  fi
done

echo ""
echo "========================"