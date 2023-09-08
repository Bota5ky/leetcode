package question_51_to_100

// https://leetcode-cn.com/problems/gray-code/
func grayCode(n int) []int {
	res := []int{0}
	for i := 1; i <= n; i++ {
		k := len(res)
		for j := k - 1; j >= 0; j-- {
			res = append(res, res[j]+1<<(i-1))
		}
	}
	return res
}
