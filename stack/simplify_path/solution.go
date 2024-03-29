package simplify_path

import "strings"

// 71. 简化路径 https://leetcode.cn/problems/simplify-path/
func simplifyPath(path string) string {
	words := strings.Split(path, "/")
	var ret []string
	for i := 0; i < len(words); i++ {
		if words[i] == "." || words[i] == "" {
			continue
		} else if words[i] == ".." {
			if len(ret) > 0 {
				ret = ret[:len(ret)-1]
			}
		} else {
			ret = append(ret, words[i])
		}
	}
	return "/" + strings.Join(ret, "/")
}
