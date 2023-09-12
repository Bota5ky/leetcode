package serialize_and_deserialize_binary_tree

import (
	. "leetcode/model"
	"strconv"
	"strings"
)

// Codec 297. 二叉树的序列化与反序列化 https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/
type Codec struct{}

func Constructor() Codec { return Codec{} }

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	var res string
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		cnt := 0
		k := len(stack)
		for i := 0; i < k; i++ {
			if stack[i] == nil {
				res += "null,"
			} else {
				stack = append(stack, stack[i].Left, stack[i].Right)
				if stack[i].Left != nil {
					cnt++
				}
				if stack[i].Right != nil {
					cnt++
				}
				res += strconv.Itoa(stack[i].Val) + ","
			}
		}
		if cnt == 0 {
			break
		}
		stack = stack[k:]
	}
	return res[:len(res)-1]
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if data == "null" {
		return nil
	}
	dataList := strings.Split(data, ",")
	num, _ := strconv.Atoi(dataList[0])
	dataList = dataList[1:]
	root := &TreeNode{Val: num}
	stack := []*TreeNode{root}
	for len(dataList) > 0 {
		k := len(stack)
		for i := 0; i < k; i++ {
			if dataList[0] != "null" {
				num, _ := strconv.Atoi(dataList[0])
				node := &TreeNode{Val: num}
				stack[i].Left = node
				stack = append(stack, node)
			}
			if dataList[1] != "null" {
				num, _ := strconv.Atoi(dataList[1])
				node := &TreeNode{Val: num}
				stack[i].Right = node
				stack = append(stack, node)
			}
			dataList = dataList[2:]
		}
		stack = stack[k:]
	}
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
