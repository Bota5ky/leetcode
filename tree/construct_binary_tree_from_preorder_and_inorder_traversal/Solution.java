package tree.construct_binary_tree_from_preorder_and_inorder_traversal;

import _model.TreeNode;

import java.util.Arrays;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/">105. 从前序与中序遍历序列构造二叉树</a>
 * @since 2024-02-21 19:42
 */
class Solution {
    public TreeNode buildTree(int[] preorder, int[] inorder) {
        if (preorder.length == 0) {
            return null;
        }
        TreeNode root = new TreeNode(preorder[0]);
        int idx = 0;
        for (int i = 0; i < inorder.length; i++) {
            if (inorder[i] == preorder[0]) {
                idx = i;
                break;
            }
        }
        root.left = buildTree(Arrays.copyOfRange(preorder, 1,idx + 1),
                Arrays.copyOf(inorder, idx));
        root.right = buildTree(Arrays.copyOfRange(preorder, idx + 1, preorder.length),
                Arrays.copyOfRange(inorder, idx + 1, inorder.length));
        return root;
    }
}
