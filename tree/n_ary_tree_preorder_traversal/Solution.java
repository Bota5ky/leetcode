package tree.n_ary_tree_preorder_traversal;

import _model.Node;

import java.util.ArrayList;
import java.util.List;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/n-ary-tree-preorder-traversal/">589. N 叉树的前序遍历</a>
 * @since 2024-02-18 19:45
 */
class Solution {
    public List<Integer> preorder(Node root) {
        List<Integer> res = new ArrayList<>();
        if (root == null) {
            return res;
        }
        res.add(root.val);
        for (Node child : root.children) {
            res.addAll(preorder(child));
        }
        return res;
    }
}
