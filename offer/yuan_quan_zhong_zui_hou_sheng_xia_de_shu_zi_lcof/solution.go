package yuan_quan_zhong_zui_hou_sheng_xia_de_shu_zi_lcof

// 剑指 Offer 62. 圆圈中最后剩下的数字 https://leetcode.cn/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/
func lastRemaining(n int, m int) int {
	last := 0
	for i := 2; i <= n; i++ {
		last = (last + m) % i
	}
	return last
}
