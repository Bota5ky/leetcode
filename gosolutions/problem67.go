package gosolutions

// https://leetcode-cn.com/problems/add-binary/
func addBinary(a string, b string) string {
	plus := 0
	var res string
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || plus > 0; {
		sum := plus
		if i >= 0 && a[i] == '1' {
			sum++
		}
		if j >= 0 && b[j] == '1' {
			sum++
		}
		switch sum {
		case 0:
			plus = 0
			res = "0" + res
		case 1:
			plus = 0
			res = "1" + res
		case 2:
			plus = 1
			res = "0" + res
		case 3:
			plus = 1
			res = "1" + res
		}
		i--
		j--
	}
	return res
}
