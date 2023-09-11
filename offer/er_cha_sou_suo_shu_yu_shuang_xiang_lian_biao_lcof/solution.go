package offer

// 未测试
// https://leetcode.cn/problems/er-cha-sou-suo-shu-yu-shuang-xiang-lian-biao-lcof/
var prenode *TreeNode  //比当前root节点小的节点
var realHead *TreeNode //定义链表头部的结点，最小值
// 中序递归遍历修改链表指针即可实现
func treeToDoublyList(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	treeToDoublyList(root.Left) //左
	if prenode == nil {         //根
		prenode = root
		realHead = root
	} else {
		prenode.Right = root
		root.Left = prenode
		prenode = root
	}
	treeToDoublyList(root.Right) //右
	return realHead
}
