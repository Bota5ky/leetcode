package leetcode

//https://leetcode-cn.com/problems/prime-number-of-set-bits-in-binary-representation/
//[2, 3, 5, 7, 11, 13, 17, 19]
func countPrimeSetBits(L int, R int) int {
	ret :=0
	for i := L; i <= R; i++ {
		num := i
		cnt := 0
		for num > 0 {
			if num&1 == 1 {
				cnt++
			}
			num >>=1
		}
		if cnt==2 || cnt==3 || cnt==5 || cnt==7 || cnt==11 || cnt==13 || cnt==17 || cnt==19 {ret++}
	}
	return ret
}
