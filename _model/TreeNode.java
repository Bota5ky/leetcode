package _model;

import java.util.LinkedList;
import java.util.Queue;

public class TreeNode {
    public int val;
    public TreeNode left;
    public TreeNode right;

    public TreeNode() {
    }

    public TreeNode(int val) {
        this.val = val;
    }

    public TreeNode(int val, TreeNode left, TreeNode right) {
        this.val = val;
        this.left = left;
        this.right = right;
    }

    public static TreeNode build(Integer[] nums) {
        return build(nums, 0);
    }

    public static TreeNode build(Integer[] nums, int i) {
        if (i >= nums.length || nums[i] == null) {
            return null;
        }
        TreeNode node = new TreeNode(nums[i]);
        node.left = build(nums, 2 * i + 1);
        node.right = build(nums, 2 * i + 2);
        return node;
    }

    public void print() {
        Queue<TreeNode> queue = new LinkedList<>();
        queue.offer(this);

        while (!queue.isEmpty()) {
            int levelSize = queue.size();
            for (int i = 0; i < levelSize; i++) {
                TreeNode node = queue.poll();
                if (node != null) {
                    System.out.printf("%-6d", node.val);
                    queue.offer(node.left);
                    queue.offer(node.right);
                } else {
                    System.out.print("#     ");
                }
            }
            System.out.println();
        }
    }
}
