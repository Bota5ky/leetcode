package partition_array_into_three_parts_with_equal_sum

// https://leetcode.cn/problems/partition-array-into-three-parts-with-equal-sum/
func canThreePartsEqualSum(A []int) bool {
	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
	}
	if sum%3 != 0 {
		return false
	}
	want := sum / 3
	cnt := 0
	sum = 0
	var i int
	for i = 0; i < len(A); i++ {
		sum += A[i]
		if sum == want {
			cnt++
			sum = 0
		}
		if cnt == 2 {
			i++
			break
		}
	}
	for ; i < len(A); i++ {
		sum += A[i]
		if sum == want && i == len(A)-1 {
			return true
		}
	}
	return false
}
