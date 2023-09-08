package question_151_to_200

// https://leetcode-cn.com/problems/implement-trie-prefix-tree/
type trie struct {
	isEnd bool
	node  [26]*trie
}

/** Initialize your data structure here. */
func constructor() trie {
	return trie{}
}

/** Inserts a word into the trie. */
func (t *trie) insert(word string) {
	for i := 0; i < len(word); i++ {
		if t.node[word[i]-'a'] == nil {
			temp := &trie{isEnd: i == len(word)-1}
			t.node[word[i]-'a'] = temp
		} else if t.node[word[i]-'a'].isEnd == false && i == len(word)-1 {
			t.node[word[i]-'a'].isEnd = true
		}
		t = t.node[word[i]-'a']
	}
}

/** Returns if the word is in the trie. */
func (t *trie) search(word string) bool {
	for i := 0; i < len(word); i++ {
		if t.node[word[i]-'a'] == nil {
			return false
		}
		if i == len(word)-1 && t.node[word[i]-'a'].isEnd == false {
			return false
		}
		t = t.node[word[i]-'a']
	}
	return true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (t *trie) startsWith(prefix string) bool {
	for i := 0; i < len(prefix); i++ {
		if t.node[prefix[i]-'a'] == nil {
			return false
		}
		t = t.node[prefix[i]-'a']
	}
	return true
}

/**
 * Your trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
