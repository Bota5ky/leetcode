package offer

// 剑指 Offer 66. 构建乘积数组 https://leetcode.cn/problems/gou-jian-cheng-ji-shu-zu-lcof/
func constructArr(a []int) []int {
	if len(a) == 0 {
		return []int{}
	}
	list := make([]int, len(a))
	list[0] = 1
	for i := 1; i < len(a); i++ {
		list[i] = list[i-1] * a[i-1]
	}
	res := make([]int, len(a))
	num := 1
	for i := len(a) - 1; i >= 0; i-- {
		if i == len(a)-1 {
			res[i] = list[i] * num
		} else {
			num *= a[i+1]
			res[i] = list[i] * num
		}
	}
	return res
}

// 从左到右 从右到左 各遍历一次
