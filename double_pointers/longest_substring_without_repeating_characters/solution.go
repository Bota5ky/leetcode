package zui_chang_bu_han_zhong_fu_zi_fu_de_zi_zi_fu_chuan_lcof

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	strMap := make(map[byte]int)
	strMap[s[0]] = 1 //位置在strMap[s[0]]-1
	head := 0
	rear := 1
	maximum := 1
	for rear < len(s) {
		if strMap[s[rear]] > 0 {
			if head < strMap[s[rear]] {
				head = strMap[s[rear]]
			}
			strMap[s[rear]] = rear + 1
		} else {
			strMap[s[rear]] = rear + 1
		}
		if rear-head+1 > maximum {
			maximum = rear - head + 1
		}
		rear++
	}
	return maximum
}
