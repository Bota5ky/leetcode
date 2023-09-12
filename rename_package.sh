#!/bin/bash

for file in common/**/**/*.go ; do
  dir_name=$(echo "$file" | awk -F'/' '{print $(NF-1)}')
  sed -i '' "1s/.*/package $dir_name/" $file
done
