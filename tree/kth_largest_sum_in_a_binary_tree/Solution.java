package tree.kth_largest_sum_in_a_binary_tree;

import _model.TreeNode;

import java.util.Comparator;
import java.util.LinkedList;
import java.util.PriorityQueue;
import java.util.Queue;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/kth-largest-sum-in-a-binary-tree/">2583. 二叉树中的第 K 大层和</a>
 * @since 2024-02-23 20:15
 */
class Solution {
    public long kthLargestLevelSum(TreeNode root, int k) {
        PriorityQueue<Long> q = new PriorityQueue<>(Comparator.reverseOrder());
        Queue<TreeNode> nodes = new LinkedList<>();
        nodes.add(root);
        int size = 1;
        while (!nodes.isEmpty()) {
            long sum = 0;
            while (size > 0) {
                size--;
                TreeNode node = nodes.poll();
                sum += node.val;
                if (node.left != null) {
                    nodes.add(node.left);
                }
                if (node.right != null) {
                    nodes.add(node.right);
                }
            }
            q.add(sum);
            size = nodes.size();
        }
        if (q.size() < k) {
            return -1;
        }
        for (int i = 0; i < k - 1; i++) {
            q.poll();
        }
        return q.poll();
    }
}
