package add_without_plus_lcci

// 面试题 17.01. 不用加号的加法 https://leetcode.cn/problems/add-without-plus-lcci/
// LCR 190. 加密运算 https://leetcode.cn/problems/bu-yong-jia-jian-cheng-chu-zuo-jia-fa-lcof/
func add(a int, b int) int {
	for b != 0 {
		tmp := a ^ b     //不计进位的相加结果
		b = (a & b) << 1 //进位的部分
		a = tmp
	}
	return a
}
