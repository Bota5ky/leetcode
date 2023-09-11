#!/bin/bash

find "." -type f -exec sed -i '' -e "s/leetcode-cn\.com/leetcode.cn/g" {} +
