package tree.construct_binary_tree_from_inorder_and_postorder_traversal;

import _model.TreeNode;

import java.util.Arrays;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal/">106. 从中序与后序遍历序列构造二叉树</a>
 * @since 2023-10-28 16:26
 */
class Solution {
    public TreeNode buildTree(int[] inorder, int[] postorder) {
        for (int i = 0; i < inorder.length; i++) {
            int rootNodeVal = postorder[postorder.length - 1];
            if (inorder[i] == rootNodeVal) {
                TreeNode node = new TreeNode(rootNodeVal);
                int[] leftNodes = Arrays.copyOfRange(inorder, 0, i);
                int[] rightNodes = Arrays.copyOfRange(inorder, i + 1, inorder.length);
                node.left = buildTree(leftNodes, Arrays.copyOfRange(postorder, 0, i));
                node.right = buildTree(rightNodes, Arrays.copyOfRange(postorder, i, postorder.length - 1));
                return node;
            }
        }
        return null;
    }
}
