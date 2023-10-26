package utf_8_validation

// 393. UTF-8 编码验证 https://leetcode.cn/problems/utf-8-validation/
func validUtf8(data []int) bool {
	for i := 0; i < len(data); {
		k := figure(data[i] % 256)
		i++
		if k < 0 {
			return false
		}
		for i < len(data) && k > 0 {
			if figure2(data[i] % 256) {
				return false
			}
			i++
			k--
		}
		if k > 0 {
			return false
		}
	}
	return true
}

func figure(num int) int {
	num >>= 3
	if num == 30 {
		return 3
	}
	num >>= 1
	if num == 14 {
		return 2
	}
	num >>= 1
	if num == 6 {
		return 1
	}
	num >>= 2
	if num == 0 {
		return 0
	}
	return -1
}

func figure2(num int) bool {
	num >>= 6
	if num == 2 {
		return false
	}
	return true
}
