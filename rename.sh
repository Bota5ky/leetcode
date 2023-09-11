#!/bin/bash

# 定义目标目录
directory="interview"

# 递归查找并重命名*.go文件为solution.go
find "$directory" -type f -name "*.go" -execdir mv {} solution.go \;

echo "重命名完成！"
