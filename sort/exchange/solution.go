package diao_zheng_shu_zu_shun_xu_shi_qi_shu_wei_yu_ou_shu_qian_mian_lcof

// LCR 139. 训练计划 I https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
func exchange(nums []int) []int {
	i := 0
	j := len(nums) - 1
	for i < j {
		for nums[i]%2 == 1 && i < j {
			i++
		}
		for nums[j]%2 == 0 && i < j {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
	return nums
}
