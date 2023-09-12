package generate_parentheses

// 22. 括号生成 https://leetcode.cn/problems/generate-parentheses/
func generateParenthesis(n int) []string {
	var res []string
	var str string
	dfs(n, 0, 0, str, &res)
	return res
}

func dfs(n, left, right int, str string, res *[]string) {
	if left > n || left < right {
		return
	}
	if left == n && right == n {
		*res = append(*res, str)
	}
	dfs(n, left+1, right, str+"(", res)
	dfs(n, left, right+1, str+")", res)
}
