package question_151_to_200

// https://leetcode-cn.com/problems/word-search-ii/submissions/
func findWords(board [][]byte, words []string) []string {
	//1.构建前缀树
	root := &Trie{}
	for i := 0; i < len(words); i++ {
		head := root
		word := words[i]
		for j := 0; j < len(word); j++ {
			if head.Next[word[j]-'a'] == nil {
				var str string
				if j == len(word)-1 {
					str = word
				}
				temp := &Trie{EndStr: str}
				head.Next[word[j]-'a'] = temp
			} else if j == len(word)-1 && head.Next[word[j]-'a'].EndStr == "" {
				head.Next[word[j]-'a'].EndStr = word
			}
			head = head.Next[word[j]-'a']
		}
	}
	//2.遍历
	var res []string
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfsTrie(board, i, j, root, &res)
		}
	}
	return res
}
func dfsTrie(board [][]byte, i, j int, root *Trie, res *[]string) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[i]) {
		return
	}
	if root.Next[board[i][j]-'a'] == nil {
		return
	}
	if root.Next[board[i][j]-'a'].EndStr != "" {
		*res = append(*res, root.Next[board[i][j]-'a'].EndStr)
		root.Next[board[i][j]-'a'].EndStr = "" //剪枝，防止遍历到重复的单词
	}
	tempStr := board[i][j] //防止重复经过
	board[i][j] = 'z' + 1
	dfsTrie(board, i-1, j, root.Next[tempStr-'a'], res)
	dfsTrie(board, i+1, j, root.Next[tempStr-'a'], res)
	dfsTrie(board, i, j-1, root.Next[tempStr-'a'], res)
	dfsTrie(board, i, j+1, root.Next[tempStr-'a'], res)
	board[i][j] = tempStr
}

// Trie 前缀树
type Trie struct {
	EndStr string //不为空即是某一单词的结尾
	Next   [27]*Trie
}
