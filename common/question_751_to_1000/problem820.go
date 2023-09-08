package question_751_to_1000

// https://leetcode-cn.com/problems/short-encoding-of-words/
func minimumLengthEncoding(words []string) int {
	cnt := 0
	root := &trie2{}
	for i := 0; i < len(words); i++ {
		word, cur := words[i], root
		for j := len(word) - 1; j >= 0; j-- {
			if cur.next[word[j]-'a'] == nil {
				newnode := &trie2{}
				cur.next[word[j]-'a'] = newnode
				if j == 0 {
					newnode.length = len(word)
					cnt += len(word) + 1
				}
			} else if j != 0 && cur.next[word[j]-'a'].length > 0 {
				cnt -= cur.next[word[j]-'a'].length + 1
				cur.next[word[j]-'a'].length = 0
			}
			cur = cur.next[word[j]-'a']
		}
	}
	return cnt
}

// 倒序字典树
type trie2 struct {
	length int
	next   [26]*trie2
}
