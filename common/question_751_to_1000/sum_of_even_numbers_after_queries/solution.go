package sum_of_even_numbers_after_queries

// https://leetcode.cn/problems/sum-of-even-numbers-after-queries/
func sumEvenAfterQueries(A []int, queries [][]int) []int {
	sum := 0
	for i := 0; i < len(A); i++ {
		if A[i]%2 == 0 {
			sum += A[i]
		}
	}
	ret := make([]int, 0)
	for i := 0; i < len(queries); i++ {
		if A[queries[i][1]]%2 != 0 && queries[i][0]%2 != 0 {
			sum += A[queries[i][1]] + queries[i][0]
		} else if A[queries[i][1]]%2 == 0 && queries[i][0]%2 == 0 {
			sum += queries[i][0]
		} else if A[queries[i][1]]%2 == 0 && queries[i][0]%2 != 0 {
			sum -= A[queries[i][1]]
		}
		A[queries[i][1]] += queries[i][0]
		ret = append(ret, sum)
	}
	return ret
}
