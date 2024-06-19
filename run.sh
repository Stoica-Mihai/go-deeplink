#!/bin/sh
set -e
compileDir=$(mktemp)

go build -o "$compileDir" ./app

exec "$compileDir"