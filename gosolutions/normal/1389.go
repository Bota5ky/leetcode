package leetcode

//https://leetcode-cn.com/problems/create-target-array-in-the-given-order/
func createTargetArray(nums []int, index []int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		res = append(res, nums[i])
		for j := len(res) - 2; j >= index[i]; j-- {
			res[j], res[j+1] = res[j+1], res[j]
		}
	}
	return res
}
