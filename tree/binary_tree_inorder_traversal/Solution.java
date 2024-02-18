package tree.binary_tree_inorder_traversal;

import _model.TreeNode;

import java.util.ArrayList;
import java.util.List;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/binary-tree-inorder-traversal/">94. 二叉树的中序遍历</a>
 * @since 2024-02-18 22:57
 */
class Solution {
    public List<Integer> inorderTraversal(TreeNode root) {
        List<Integer> res = new ArrayList<>();
        if (root == null) {
            return res;
        }
        res.addAll(inorderTraversal(root.left));
        res.add(root.val);
        res.addAll(inorderTraversal(root.right));
        return res;
    }
}
