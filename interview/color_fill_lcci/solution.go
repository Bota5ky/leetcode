package color_fill_lcci

// 面试题 08.10. 颜色填充 https://leetcode.cn/problems/color-fill-lcci/
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	fillColor(image, sr, sc, image[sr][sc], newColor)
	return image
}

func fillColor(image [][]int, x int, y int, oldColor int, newColor int) {
	if oldColor == newColor {
		return
	}
	image[x][y] = newColor
	if x > 0 && image[x-1][y] == oldColor {
		image[x-1][y] = newColor
		fillColor(image, x-1, y, oldColor, newColor)
	}
	if x < len(image)-1 && image[x+1][y] == oldColor {
		image[x+1][y] = newColor
		fillColor(image, x+1, y, oldColor, newColor)
	}
	if y > 0 && image[x][y-1] == oldColor {
		image[x][y-1] = newColor
		fillColor(image, x, y-1, oldColor, newColor)
	}
	if y < len(image[x])-1 && image[x][y+1] == oldColor {
		image[x][y+1] = newColor
		fillColor(image, x, y+1, oldColor, newColor)
	}
}
