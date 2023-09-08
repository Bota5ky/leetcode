package common

import (
	"strconv"
	"strings"
)

// 层序
type codec struct{}

func constructor4() codec { return codec{} }

// Serializes a tree to a single string.层序遍历
func (c *codec) serialize(root *TreeNode) string {
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
func (c *codec) deserialize(data string) *TreeNode {
	if data == "null" {
		return nil
	}
	datas := strings.Split(data, ",")
	num, _ := strconv.Atoi(datas[0])
	datas = datas[1:]
	root := &TreeNode{Val: num}
	stack := []*TreeNode{root}
	for len(datas) > 0 {
		k := len(stack)
		for i := 0; i < k; i++ {
			if datas[0] != "null" {
				num, _ := strconv.Atoi(datas[0])
				node := &TreeNode{Val: num}
				stack[i].Left = node
				stack = append(stack, node)
			}
			if datas[1] != "null" {
				num, _ := strconv.Atoi(datas[1])
				node := &TreeNode{Val: num}
				stack[i].Right = node
				stack = append(stack, node)
			}
			datas = datas[2:]
		}
		stack = stack[k:]
	}
	return root
}

/**
 * Your codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */

//前序
/* type codec struct {
	sb    strings.Builder
	index int
}

func Constructor() codec {
	return codec{}
}

// Serializes a tree to a single string.
func (this *codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	this.sb = strings.Builder{}
	this.dfsSerialize(root)
	str := this.sb.String()
	return str[:len(str)-1]
}

func (this *codec) dfsSerialize(root *TreeNode) {
	if root == nil {
		this.sb.Write([]byte("n,"))
	} else {
		this.sb.Write([]byte(strconv.Itoa(root.Val) + ","))
		this.dfsSerialize(root.Left)
		this.dfsSerialize(root.Right)
	}
}

// Deserializes your encoded data to tree.
func (this *codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}

	strs := strings.Split(data, ",")
	this.index = 0
	return this.dfsDeserialize(strs)
}

func (this *codec) dfsDeserialize(strs []string) *TreeNode {
	if this.index >= len(strs) || strs[this.index] == "n" {
		this.index++
		return nil
	}

	val, _ := strconv.Atoi(strs[this.index])
	root := &TreeNode{Val: val}
	this.index++
	root.Left = this.dfsDeserialize(strs)
	root.Right = this.dfsDeserialize(strs)

	return root
} */
