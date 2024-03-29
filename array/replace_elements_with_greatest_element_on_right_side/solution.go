package replace_elements_with_greatest_element_on_right_side

// 1299. 将每个元素替换为右侧最大元素 https://leetcode.cn/problems/replace-elements-with-greatest-element-on-right-side/
func replaceElements(arr []int) []int {
	ret := make([]int, len(arr))
	ret[len(ret)-1] = -1
	for i := len(ret) - 2; i >= 0; i++ {
		ret[i] = max(ret[i+1], arr[i+1])
	}
	return ret
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
