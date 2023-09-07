package gosolutions

import (
	"strconv"
	"strings"
)

// https://leetcode-cn.com/problems/serialize-and-deserialize-bst/
type codec1 struct {
	sb  strings.Builder
	idx int
}

func constructor5() codec1 {
	return codec1{}
}
func (c *codec1) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	c.sb = strings.Builder{}
	c.dfsSerialize(root)
	str := c.sb.String()
	return str[:len(str)-1]
}
func (c *codec1) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	nums := strings.Split(data, ",")
	return c.dfsDeserialize(nums)
}
func (c *codec1) dfsDeserialize(nums []string) *TreeNode {
	index := c.idx
	if index > len(nums) || nums[index] == "null" {
		c.idx++
		return nil
	}
	val, _ := strconv.Atoi(nums[index])
	root := &TreeNode{Val: val}
	c.idx++
	root.Left = c.dfsDeserialize(nums)
	root.Right = c.dfsDeserialize(nums)
	return root
}
func (c *codec1) dfsSerialize(root *TreeNode) {
	if root == nil {
		c.sb.WriteString("null,")
		return
	}
	c.sb.WriteString(strconv.Itoa(root.Val) + ",")
	c.dfsSerialize(root.Left)
	c.dfsSerialize(root.Right)
}
