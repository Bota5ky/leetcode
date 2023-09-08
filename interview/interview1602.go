package interview

// https://leetcode-cn.com/problems/words-frequency-lcci/
type wordsFrequency struct {
	count map[string]int
}

func constructor(book []string) wordsFrequency {
	t := wordsFrequency{
		count: make(map[string]int, len(book)),
	}
	for _, v := range book {
		t.count[v]++
	}
	return t
}

func (t *wordsFrequency) get(word string) int {
	return t.count[word]
}

/**
 * Your WordsFrequency object will be instantiated and called as such:
 * obj := Constructor(book);
 * param_1 := obj.Get(word);
 */
