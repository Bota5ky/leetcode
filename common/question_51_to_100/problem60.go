package question_51_to_100

// https://leetcode-cn.com/problems/permutation-sequence/
func getPermutation(n int, k int) string {
	//1.剩下n-1位数共有(n-1)!次排列方式  首位为nums[k/(n-1)!]
	var res string
	nums, fac := make([]int, n), 1
	for i := 0; i < n; i++ {
		nums[i] = i + 1
		fac *= i + 1
	}
	k-- // 0 ~ n!-1
	for ; len(nums) > 0; n-- {
		fac /= n
		pos := k / fac
		k %= fac
		res += string('0' + byte(nums[pos]))
		nums = append(nums[:pos], nums[pos+1:]...)
	}
	return res
}
