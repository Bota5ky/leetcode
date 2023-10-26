package implement_trie_prefix_tree

// Trie 208. 实现 Trie (前缀树) https://leetcode.cn/problems/implement-trie-prefix-tree/
type Trie struct {
	isEnd bool
	node  [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the Trie. */
func (t *Trie) insert(word string) {
	for i := 0; i < len(word); i++ {
		if t.node[word[i]-'a'] == nil {
			temp := &Trie{isEnd: i == len(word)-1}
			t.node[word[i]-'a'] = temp
		} else if t.node[word[i]-'a'].isEnd == false && i == len(word)-1 {
			t.node[word[i]-'a'].isEnd = true
		}
		t = t.node[word[i]-'a']
	}
}

/** Returns if the word is in the Trie. */
func (t *Trie) search(word string) bool {
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

/** Returns if there is any word in the Trie that starts with the given prefix. */
func (t *Trie) startsWith(prefix string) bool {
	for i := 0; i < len(prefix); i++ {
		if t.node[prefix[i]-'a'] == nil {
			return false
		}
		t = t.node[prefix[i]-'a']
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
