package words_frequency_lcci

// WordsFrequency 面试题 16.02. 单词频率 https://leetcode.cn/problems/words-frequency-lcci/
type WordsFrequency struct {
	count map[string]int
}

func Constructor(book []string) WordsFrequency {
	t := WordsFrequency{
		count: make(map[string]int, len(book)),
	}
	for _, v := range book {
		t.count[v]++
	}
	return t
}

func (t *WordsFrequency) get(word string) int {
	return t.count[word]
}

/**
 * Your WordsFrequency object will be instantiated and called as such:
 * obj := Constructor(book);
 * param_1 := obj.Get(word);
 */
