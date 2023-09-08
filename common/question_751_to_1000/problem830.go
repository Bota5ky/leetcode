package question_751_to_1000

// https://leetcode-cn.com/problems/positions-of-large-groups/
func largeGroupPositions(S string) [][]int {
	var ret [][]int
	fast, slow := 0, 0
	for i := 0; i < len(S); i++ {
		fast++
		for fast < len(S) && S[fast] == S[slow] {
			fast++
		}
		if fast-slow >= 3 {
			ret = append(ret, []int{slow, fast - 1})
		}
		slow = fast
	}
	return ret
}
