package count_number_of_teams

// 1395. 统计作战单位数 https://leetcode.cn/problems/count-number-of-teams/
func numTeams(rating []int) int {
	cnt := 0
	for i := 0; i < len(rating)-2; i++ {
		for j := i + 1; j < len(rating)-1; j++ {
			for k := j + 1; k < len(rating); k++ {
				if rating[i] < rating[j] && rating[j] < rating[k] || rating[i] > rating[j] && rating[j] > rating[k] {
					cnt++
				}
			}
		}
	}
	return cnt
}
