package happy_number

// 202. 快乐数 https://leetcode.cn/problems/happy-number/
func isHappy(n int) bool {
	m := make(map[int]bool)
	for n != 1 {
		if m[n] {
			return false
		}
		m[n] = true
		n = deal(n)
	}
	return true
}

func deal(n int) int {
	res := 0
	for n > 0 {
		res += (n % 10) * (n % 10)
		n /= 10
	}
	return res
}
