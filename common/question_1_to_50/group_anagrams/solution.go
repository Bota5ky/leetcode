package group_anagrams

// 49. 字母异位词分组 https://leetcode.cn/problems/group-anagrams/
func groupAnagrams(strs []string) [][]string {
	var ret [][]string
	cnt := 0
	m := make(map[[26]int]int, 0)
	for _, str := range strs {
		key := count(str)
		_, ok := m[key]
		if !ok {
			m[key] = cnt
			ret = append(ret, []string{str})
			cnt++
		} else {
			ret[m[key]] = append(ret[m[key]], str)
		}
	}
	return ret
}

func count(str string) [26]int {
	var ret [26]int
	for _, v := range str {
		ret[v-'a']++
	}
	return ret
}
