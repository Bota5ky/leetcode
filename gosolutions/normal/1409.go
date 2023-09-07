package leetcode

//https://leetcode-cn.com/problems/queries-on-a-permutation-with-key/
func processQueries(queries []int, m int) []int {
	var res []int
	p := make(map[int]int, m)
	num := make([]int, m)
	for i := 0; i < m; i++ {
		p[i+1] = i //位置
		num[i] = i + 1
	}
	for i := 0; i < len(queries); i++ {
		res = append(res, p[queries[i]])
		temp := num[p[queries[i]]]
		for j := p[queries[i]] - 1; j >= 0; j-- {
			p[num[j]]++
			num[j+1] = num[j]
		}
		num[0] = temp
		p[temp] = 0
	}
	return res
}
