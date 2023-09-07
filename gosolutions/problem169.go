package gosolutions

// 和offer39相同，这里采用 Boyer-Moore 投票算法
// https://leetcode-cn.com/problems/majority-element/
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
