package er_jin_zhi_zhong_1de_ge_shu_lcof

// LCR 133. 位 1 的个数 https://leetcode.cn/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/
func hammingWeight(num uint32) int {
	var cnt uint32
	for num != 0 {
		cnt += num & 1
		num >>= 1
	}
	return int(cnt)
}
