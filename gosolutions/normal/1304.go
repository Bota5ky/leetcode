package leetcode

//https://leetcode-cn.com/problems/find-n-unique-integers-sum-up-to-zero/
func sumZero(n int) []int {
    var ret []int
    for i:=-n/2; i<=n/2; i++ {
        if i==0 && n%2==0 {continue}
        ret=append(ret,i)
    }
    return ret
}
