#!/bin/bash

find "." -type f -name "*.go" -execdir mv {} solution.go \;

echo "重命名完成！"
