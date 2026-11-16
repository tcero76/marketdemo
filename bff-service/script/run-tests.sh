#!/bin/sh

set -a
. ./test.env
set +a

exec go test -count=1 -v \
  -run "$TEST_PATTERN" \
  "./$TEST_PACKAGE/..."