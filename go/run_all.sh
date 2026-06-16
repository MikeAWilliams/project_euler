#!/usr/bin/env bash
set -euo pipefail

for dir in p*/; do
    echo "=== ${dir%/} ==="
    (cd "$dir" && go run .) && echo
done
