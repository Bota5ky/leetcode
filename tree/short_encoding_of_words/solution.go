package short_encoding_of_words

// 820. 单词的压缩编码 https://leetcode.cn/problems/short-encoding-of-words/
func minimumLengthEncoding(words []string) int {
	cnt := 0
	root := &trie{}
	for i := 0; i < len(words); i++ {
		word, cur := words[i], root
		for j := len(word) - 1; j >= 0; j-- {
			if cur.next[word[j]-'a'] == nil {
				newNode := &trie{}
				cur.next[word[j]-'a'] = newNode
				if j == 0 {
					newNode.length = len(word)
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
type trie struct {
	length int
	next   [26]*trie
}
