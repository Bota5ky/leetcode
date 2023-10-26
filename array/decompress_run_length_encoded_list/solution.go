package decompress_run_length_encoded_list

// 1313. 解压缩编码列表 https://leetcode.cn/problems/decompress-run-length-encoded-list/
func decompressRLElist(nums []int) []int {
	var ret []int
	if len(nums)%2 == 1 {
		return ret
	}
	for i := 0; i < len(nums); i += 2 {
		for j := 0; j < nums[i]; j++ {
			ret = append(ret, nums[i+1])
		}
	}
	return ret
}
