package majority_element

// 169. 多数元素 https://leetcode.cn/problems/majority-element/
func majorityElement(nums []int) int {
	temp, cnt := 0, 0
	for _, v := range nums {
		if cnt == 0 || v == temp {
			temp = v
			cnt++
		} else {
			cnt--
		}
	}
	return temp
}
