package question_751_to_1000

import "math"

// https://leetcode.cn/problems/largest-triangle-area/
func largestTriangleArea(points [][]int) float64 {
	var area float64
	for i := 0; i < len(points)-2; i++ {
		for j := i + 1; j < len(points)-1; j++ {
			for k := j + 1; k < len(points); k++ {
				x1 := points[i][0]
				y1 := points[i][1]
				x2 := points[j][0]
				y2 := points[j][1]
				x3 := points[k][0]
				y3 := points[k][1]
				area = math.Max(area, 0.5*math.Abs(float64(x1*y2+x2*y3+x3*y1-x1*y3-x2*y1-x3*y2)))
			}
		}
	}
	return area
}
