package gas_station

// 134. 加油站 https://leetcode.cn/problems/gas-station/
func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); i++ {
		if gas[i] < cost[i] {
			continue
		}
		store := gas[i] - cost[i]
		j := i + 1
		if j == len(gas) {
			j = 0
		}
		for ; j != i && store+gas[j] >= cost[j]; j++ {
			store = store + gas[j] - cost[j]
			if j == len(gas)-1 {
				j = -1
			}
		}
		if j == i {
			return i
		}
	}
	return -1
}
