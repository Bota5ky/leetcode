package tree.n_ary_tree_level_order_traversal;

import _model.Node;

import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;
import java.util.Queue;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/n-ary-tree-level-order-traversal/">429. N 叉树的层序遍历</a>
 * @since 2024-02-17 19:12
 */
class Solution {
    public List<List<Integer>> levelOrder(Node root) {
        List<List<Integer>> res = new ArrayList<>();
        if (root == null) {
            return res;
        }
        Queue<Node> queue = new LinkedList<>();
        queue.add(root);
        int cnt = 1;
        while (!queue.isEmpty()) {
            List<Integer> layer = new ArrayList<>();
            while (cnt > 0) {
                cnt--;
                Node node = queue.poll();
                layer.add(node.val);
                if (node.children != null) {
                    for (Node child : node.children) {
                        if (child != null) {
                            queue.add(child);
                        }
                    }
                }
            }
            cnt = queue.size();
            res.add(layer);
        }
        return res;
    }
}
