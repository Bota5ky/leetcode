package leetcode

//https://leetcode-cn.com/problems/count-square-submatrices-with-all-ones/
func countSquares(matrix [][]int) int {
	cnt:=0
    for i:= 0; i<len(matrix); i++ {
		for j:= 0; j<len(matrix[i]); j++ {
			if matrix[i][j]==0 {continue}
			cnt+=cntSquare(matrix,i,j)+1
		}
	}
	return cnt
}

func cntSquare(matrix [][]int,a int ,b int) int {
	cnt:=0
	LOOP:
	for i:=1;a+i<len(matrix);i++ {
		for j:=b;j<=b+i;j++ {
			if matrix[a+i][j]==0 {break LOOP}
		}
		for j:=a;j<a+i;j++ {
			if matrix[a+j][b]==0 {break LOOP}
		}
		cnt++
	}
	return cnt
}