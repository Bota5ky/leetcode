package common

import "strings"

// https://leetcode-cn.com/problems/license-key-formatting/
func licenseKeyFormatting(S string, K int) string {
	S = strings.ToUpper(S)
	S = strings.Replace(S, "-", "", -1)
	if len(S) < K {
		return S
	}
	first := len(S) % K
	plus := 0
	if first > 0 {
		S = S[:first] + "-" + S[first:]
		plus = 1
	}
	for i := first + plus + K; i < len(S); i += K {
		S = S[:i] + "-" + S[i:]
		i++
	}
	return S
}
