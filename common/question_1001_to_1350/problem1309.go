package question_1001_to_1350

// https://leetcode-cn.com/problems/decrypt-string-from-alphabet-to-integer-mapping/
// jklmnopqrs,tuvwxyz
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
