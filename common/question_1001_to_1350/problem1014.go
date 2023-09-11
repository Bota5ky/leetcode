package question_1001_to_1350

// https://leetcode.cn/problems/best-sightseeing-pair/
func maxScoreSightseeingPair(A []int) int {
	B := make([]int, len(A))
	for c, v := range A {
		B[c] = v - c //A[j]-j
		A[c] = v + c //A[i]+i
	}
	for i := len(B) - 2; i >= 0; i-- {
		if B[i] < B[i+1] {
			B[i] = B[i+1]
		}
	}
	for i := 1; i < len(A); i++ {
		if A[i] < A[i-1] {
			A[i] = A[i-1]
		}
	}
	max := 0
	for i := 0; i < len(A)-1; i++ {
		if A[i]+B[i+1] > max {
			max = A[i] + B[i+1]
		}
	}
	return max
}
