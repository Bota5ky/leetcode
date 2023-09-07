package gosolutions

// https://leetcode-cn.com/problems/route-between-nodes-lcci/
func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	//[0] {1,2,3,...}
	g := make([][]int, n)
	visited := make(map[int]bool, n)
	for _, v := range graph {
		g[v[0]] = append(g[v[0]], v[1])
	}
	return dfs2(g, start, target, visited)
}

func dfs2(g [][]int, start int, target int, visited map[int]bool) bool {
	visited[start] = true
	for i := 0; i < len(g[start]); i++ {
		if g[start][i] == target {
			return true
		}
		if visited[g[start][i]] {
			continue
		}
		if dfs2(g, g[start][i], target, visited) {
			return true
		}
	}
	return false
}
