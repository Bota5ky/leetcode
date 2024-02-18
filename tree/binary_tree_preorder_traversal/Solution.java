package tree.binary_tree_preorder_traversal;

import _model.TreeNode;

import java.util.ArrayList;
import java.util.List;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/binary-tree-preorder-traversal/">144. 二叉树的前序遍历</a>
 * @since 2024-02-18 22:49
 */
class Solution {
    public List<Integer> preorderTraversal(TreeNode root) {
        List<Integer> res = new ArrayList<>();
        if (root == null) {
            return res;
        }
        res.add(root.val);
        res.addAll(preorderTraversal(root.left));
        res.addAll(preorderTraversal(root.right));
        return res;
    }
}
