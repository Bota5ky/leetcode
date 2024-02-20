package tree.n_ary_tree_postorder_traversal;

import _model.Node;

import java.util.ArrayList;
import java.util.List;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/n-ary-tree-postorder-traversal/">590. N 叉树的后序遍历</a>
 * @since 2024-02-20 08:33
 */
class Solution {
    public List<Integer> postorder(Node root) {
        List<Integer> res = new ArrayList<>();
        if (root == null) {
            return res;
        }
        for (Node child : root.children) {
            res.addAll(postorder(child));
        }
        res.add(root.val);
        return res;
    }
}
