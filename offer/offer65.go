package offer

// https://leetcode-cn.com/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/
func add(a int, b int) int {
	for b != 0 {
		tmp := a ^ b     //不计进位的相加结果
		b = (a & b) << 1 //进位的部分
		a = tmp
	}
	return a
}
