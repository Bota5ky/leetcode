package gosolutions

import "math"

// https://leetcode-cn.com/problems/nge-tou-zi-de-dian-shu-lcof/
func twoSum2(n int) []float64 {
	//f(n,s)=f(n-1,s-1)+f(n-1,s-2)+f(n-1,s-3)+f(n-1,s-4)+f(n-1,s-5)+f(n-1,s-6)
	//f(1,1)=f(1,2)=f(1,3)=f(1,4)=f(1,5)=f(1,6)=1
	//s=n~6n之间
	total := math.Pow(6, float64(n))
	f := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = make([]int, 6*n+1)
	}
	//初始化
	for i := 1; i <= 6; i++ {
		f[1][i] = 1
	}
	//赋值
	for i := 2; i <= n; i++ {
		for j := i; j <= 6*i; j++ {
			f[i][j] += f[i-1][j-1]
			if j > 2 {
				f[i][j] += f[i-1][j-2]
			}
			if j > 3 {
				f[i][j] += f[i-1][j-3]
			}
			if j > 4 {
				f[i][j] += f[i-1][j-4]
			}
			if j > 5 {
				f[i][j] += f[i-1][j-5]
			}
			if j > 6 {
				f[i][j] += f[i-1][j-6]
			}
		}
	}
	//结果
	var res []float64
	for i := n; i <= 6*n; i++ {
		res = append(res, float64(f[n][i])/total)
	}
	return res
}
