#!/usr/bin/env bash
set -euo pipefail

for dir in p*/; do
    echo "=== ${dir%/} ==="
    go run "./$dir" && echo
done
