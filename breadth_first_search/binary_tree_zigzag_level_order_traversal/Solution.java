package breadth_first_search.binary_tree_zigzag_level_order_traversal;

import _model.TreeNode;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal/">103. 二叉树的锯齿形层序遍历</a>
 * @since 2024-02-16 23:18
 */
class Solution {
    public List<List<Integer>> zigzagLevelOrder(TreeNode root) {
        List<List<Integer>> res = new ArrayList<>();
        if (root == null) {
            return res;
        }
        List<TreeNode> nextLayer = new ArrayList<>();
        nextLayer.add(root);
        int direction = 0;
        while (!nextLayer.isEmpty()) {
            List<TreeNode> newLayer = new ArrayList<>();
            List<Integer> reviewed = new ArrayList<>();
            for (int i = 0; i <= nextLayer.size() - 1; i++) {
                TreeNode node = nextLayer.get(i);
                reviewed.add(node.val);
                if (node.left != null) {
                    newLayer.add(node.left);
                }
                if (node.right != null) {
                    newLayer.add(node.right);
                }
            }
            if (direction++ % 2 == 1) {
                Collections.reverse(reviewed);
            }
            res.add(reviewed);
            nextLayer = newLayer;
        }
        return res;
    }
}
