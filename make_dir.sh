#!/bin/bash

path="common/question_1350_to_inf"

for file in "$path"/*; do
  if [ -f "$file" ]; then
    snake_case_name=$(grep -o 'problems/[^/]*' "$file" | cut -d'/' -f2 | sed 's/-/_/g')
    if [ -n "$snake_case_name" ]; then
      mkdir -p "$path/$snake_case_name"
      mv "$file" "$path/$snake_case_name/solution.go"
    fi
  fi
done