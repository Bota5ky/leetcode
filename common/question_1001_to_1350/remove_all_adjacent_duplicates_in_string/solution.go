package remove_all_adjacent_duplicates_in_string

// https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/
func removeDuplicates(S string) string {
	for i := 0; i < len(S)-1; i++ {
		if S[i+1] == S[i] {
			S = S[:i] + S[i+2:]
			i = -1
		}
	}
	return S
}
