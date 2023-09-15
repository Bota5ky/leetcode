package course_schedule_ii

// 210. 课程表 II https://leetcode.cn/problems/course-schedule-ii/
// LCR 113. 课程表 II https://leetcode.cn/problems/QA2IGt/
func findOrder(numCourses int, prerequisites [][]int) []int {
	sons := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	for _, conn := range prerequisites {
		sons[conn[1]] = append(sons[conn[1]], conn[0])
		inDegree[conn[0]]++
	}
	var res []int
	var stack []int
	for key, val := range inDegree {
		if val == 0 {
			stack = append(stack, key) //入度为0的初始点
			res = append(res, key)
		}
	}
	for len(stack) > 0 {
		k := len(stack)
		for i := 0; i < k; i++ {
			for j := 0; j < len(sons[stack[i]]); j++ {
				inDegree[sons[stack[i]][j]]--
				if inDegree[sons[stack[i]][j]] == 0 {
					stack = append(stack, sons[stack[i]][j])
					res = append(res, sons[stack[i]][j])
				}
			}
		}
		stack = stack[k:]
	}
	if len(res) != numCourses {
		return []int{}
	}
	return res
}
