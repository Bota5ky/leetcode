package subsets

// 78. 子集 https://leetcode.cn/problems/subsets/
func subsets(nums []int) [][]int {
	res := [][]int{[]int{nums[0]}, []int{}}
	for i := 1; i < len(nums); i++ {
		k := len(res)
		for j := 0; j < k; j++ {
			temp := make([]int, len(res[j]))
			copy(temp, res[j])
			res = append(res, temp)
			res[j] = append(res[j], nums[i])
		}
	}
	return res
}
