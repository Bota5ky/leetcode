package gosolutions

// https://leetcode-cn.com/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-ii-lcof/
func singleNumber2(nums []int) int {
	res := make([]int, 32)
	for _, v := range nums {
		for i := 0; v > 0; i++ {
			res[i] += v & 1
			res[i] %= 3
			v >>= 1
		}
	}
	ret := 0
	for i := 31; i >= 0; i-- {
		ret <<= 1
		ret += res[i]
	}
	return ret
}

//当存在负数时	go语言中位运算取反符号与异或相同^
// //和137相同
// func singleNumber(nums []int) int {
//     once,twice:=0,0
//     for _,v:=range nums{
//         once=^twice &(once^v)
//         twice=^once &(twice^v)
//     }
//     return once
// }
