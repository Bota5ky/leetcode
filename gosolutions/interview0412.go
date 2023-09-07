package gosolutions

// https://leetcode-cn.com/problems/paths-with-sum-lcci/
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := oneRoot(root, sum)
	return res + pathSum(root.Left, sum) + pathSum(root.Right, sum)
}

func oneRoot(root *TreeNode, target int) int {
	cnt := 0
	dfs(root, 0, target, &cnt)
	return cnt
}

func dfs(root *TreeNode, curSum, target int, cnt *int) {
	if root == nil {
		return
	}
	curSum += root.Val
	if curSum == target {
		*cnt++
	}
	dfs(root.Left, curSum, target, cnt)
	dfs(root.Right, curSum, target, cnt)
}
