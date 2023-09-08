package question_301_to_450

import "strconv"

// https://leetcode-cn.com/problems/string-compression/
func compress(chars []byte) int {
	slow := 0
	for i := 0; i < len(chars); i++ {
		fast := slow + 1
		for chars[fast] == chars[slow] {
			fast++
		}
		if fast-slow > 1 {
			rear := chars[fast:]
			chars = append(chars[:slow], strconv.Itoa(fast-slow)...)
			chars = append(chars, rear...)
		}
		slow = fast
	}
	return len(chars)
}
