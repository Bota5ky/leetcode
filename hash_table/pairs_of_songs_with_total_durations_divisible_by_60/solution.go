package pairs_of_songs_with_total_durations_divisible_by_60

// 1010. 总持续时间可被 60 整除的歌曲 https://leetcode.cn/problems/pairs-of-songs-with-total-durations-divisible-by-60/
func numPairsDivisibleBy60(time []int) int {
	m := make(map[int]int)
	cnt := 0
	for c := range time {
		time[c] %= 60
		m[time[c]]++
	}
	cnt += (m[0] - 1) * m[0] / 2
	cnt += (m[30] - 1) * m[30] / 2
	for i := 1; i < 30; i++ {
		cnt += m[i] * m[60-i]
	}
	return cnt
}
