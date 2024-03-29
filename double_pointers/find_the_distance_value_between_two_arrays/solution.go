package find_the_distance_value_between_two_arrays

// 1385. 两个数组间的距离值 https://leetcode.cn/problems/find-the-distance-value-between-two-arrays/
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	cnt := 0
	for i := 0; i < len(arr1); i++ {
		for j := 0; j < len(arr2); j++ {
			if abs(arr1[i]-arr2[j]) <= d {
				cnt--
				break
			}
		}
		cnt++
	}
	return cnt
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
