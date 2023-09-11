package question_51_to_100

// https://leetcode.cn/problems/subsets/
func subsets(nums []int) [][]int {
	res := [][]int{[]int{nums[0]}, []int{}}
	for i := 1; i < len(nums); i++ {
		k := len(res)
		for j := 0; j < k; j++ {
			temp := make([]int, len(res[j]))
			copy(temp, res[j])
			res = append(res, temp)
			res[j] = append(res[j], nums[i])
		}
	}
	return res
}

/* 递归
func subsets(nums []int) [][]int {
    var res [][]int
    if len(nums)==0 {return res}
    if len(nums)==1 {
        return [][]int{[]int{nums[0]},[]int{}}
    }
    pre:=subsets(nums[1:])
    for i:=0;i<len(pre);i++ {
        temp:=make([]int,len(pre[i]))
        copy(temp,pre[i])
        pre[i]=append(pre[i],nums[0])
        res=append(res,temp)
        res=append(res,pre[i])
    }
    return res
}
*/
