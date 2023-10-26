package decrypt_string_from_alphabet_to_integer_mapping

// 1309. 解码字母到整数映射 https://leetcode.cn/problems/decrypt-string-from-alphabet-to-integer-mapping/
func freqAlphabets(s string) string {
	var ret []byte
	for i := 0; i < len(s); i++ {
		if i+2 < len(s) && s[i+2] == '#' {
			if s[i] == '1' {
				ret = append(ret, 'j'+s[i+1]-'0')
			} else {
				ret = append(ret, 't'+s[i+1]-'0')
			}
			i += 2
		} else {
			ret = append(ret, 'a'+s[i]-'1')
		}
	}
	return string(ret)
}
