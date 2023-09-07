package leetcode

//CheckPermutation 判定是否互为字符重排
//https://leetcode-cn.com/problems/check-permutation-lcci/
func CheckPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	cnt := make([]int, 256)
	for i := 0; i < len(s1); i++ {
		cnt[int(s1[i])]++
		cnt[int(s2[i])]--
	}
	for i := 0; i < 256; i++ {
		if cnt[i] != 0 {
			return false
		}
	}
	return true
}
