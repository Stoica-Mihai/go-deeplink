#!/bin/sh
set -e
compileDir=$(mktemp)

go build -o "$compileDir" ./backend

exec "$compileDir"