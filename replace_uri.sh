#!/bin/bash

# 定义目标目录
directory="offer"

# 递归查找并替换*.go和*.java文件中的字符串
find "$directory" -type f \( -name "*.go" -o -name "*.java" \) -exec sed -i 's/leetcode-cn\.com/leetcode.cn/g' {} +

echo "替换完成！"
