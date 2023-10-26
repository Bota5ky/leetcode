package serialize_and_deserialize_bst

import (
	. "leetcode/model"
	"strconv"
	"strings"
)

// Codec 449. 序列化和反序列化二叉搜索树 https://leetcode.cn/problems/serialize-and-deserialize-bst/
type Codec struct {
	sb  strings.Builder
	idx int
}

func Constructor() Codec {
	return Codec{}
}

func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	c.sb = strings.Builder{}
	c.dfsSerialize(root)
	str := c.sb.String()
	return str[:len(str)-1]
}

func (c *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	nums := strings.Split(data, ",")
	return c.dfsDeserialize(nums)
}

func (c *Codec) dfsDeserialize(nums []string) *TreeNode {
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

func (c *Codec) dfsSerialize(root *TreeNode) {
	if root == nil {
		c.sb.WriteString("null,")
		return
	}
	c.sb.WriteString(strconv.Itoa(root.Val) + ",")
	c.dfsSerialize(root.Left)
	c.dfsSerialize(root.Right)
}
