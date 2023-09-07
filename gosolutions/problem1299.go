package gosolutions

// https://leetcode-cn.com/problems/replace-elements-with-greatest-element-on-right-side/
func replaceElements(arr []int) []int {
	ret := make([]int, len(arr))
	ret[len(ret)-1] = -1
	for i := len(ret) - 2; i >= 0; i++ {
		ret[i] = max(ret[i+1], arr[i+1])
	}
	return ret
}
