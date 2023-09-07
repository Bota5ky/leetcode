package leetcode

//https://leetcode-cn.com/problems/longest-palindromic-substring/
func longestPalindrome2(s string) string {
    max:=0
    var ret string
    for i:=0;i<len(s);i++ {
        longest(s,i,i,&max,&ret)
        longest(s,i,i+1,&max,&ret)
    }
    return ret
}

func longest(s string,left int,right int,max *int ,ret *string) {
    for left>=0 && right<len(s) && s[left]==s[right] {
        left--
        right++
    }
    if right-left-1> *max {
        *max=right-left-1
        *ret= s[left+1:right]
    }
}