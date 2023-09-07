package leetcode

//https://leetcode-cn.com/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/
// func lastRemaining(n int, m int) int {
// 	nums := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		nums[i] = i
// 	}
// 	for i := 0; len(nums) > 1; {
// 		i += m - 1
// 		i %= len(nums)
// 		nums = append(nums[:i], nums[i+1:]...)
// 	}
// 	return nums[0]
// }

func lastRemaining(n int, m int) int {
	last := 0
	for i := 2; i <= n; i++ {
		last = (last + m) % i
	}
	return last
}
