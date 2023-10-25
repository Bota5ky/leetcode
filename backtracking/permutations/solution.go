package permutations

// 46. 全排列 https://leetcode.cn/problems/permutations/
func permute(nums []int) [][]int {
	var res [][]int
	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}
	for i := 0; i < len(nums); i++ {
		//剩下的数字组
		var temp []int
		temp = append(temp, nums[:i]...)
		temp = append(temp, nums[i+1:]...)
		//剩下的数字排列[][]int
		rest := permute(temp)
		for j := 0; j < len(rest); j++ {
			rest[j] = append(rest[j], nums[i])
			res = append(res, rest[j])
		}
	}
	return res
}
