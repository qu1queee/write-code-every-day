#!/usr/bin/env bash
set -euo pipefail

if [[ $# -ne 1 ]]; then
  echo "Usage: $0 <problem-dir-relative-to-go/>" >&2
  echo "Example: $0 problems/phase0/lc0225_stack_using_queues" >&2
  exit 1
fi

PROBLEM_DIR="go/$1"
SOLUTION="$PROBLEM_DIR/solution.go"

if [[ ! -f "$SOLUTION" ]]; then
  echo "No solution.go found at $SOLUTION" >&2
  exit 1
fi

STAMP=$(date +%Y%m%d)
BACKUP="$PROBLEM_DIR/solution_${STAMP}.go.bak"
INDEX=1
while [[ -e "$BACKUP" ]]; do
  BACKUP="$PROBLEM_DIR/solution_${STAMP}_${INDEX}.go.bak"
  INDEX=$((INDEX + 1))
done

mv "$SOLUTION" "$BACKUP"
echo "Backed up $SOLUTION -> $BACKUP"
echo "Create a fresh solution.go and re-solve from memory."
