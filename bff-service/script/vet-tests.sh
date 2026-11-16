#!/bin/sh

set -a
. ./test.env
set +a

exec go vet ./${TEST_PACKAGE:-controller}/...