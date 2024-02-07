package _model;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;
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
        Queue<Integer> queue = new LinkedList<>(Arrays.asList(nums));
        Integer first = queue.poll();
        TreeNode root = create(first);
        List<TreeNode> layer = new ArrayList<>();
        layer.add(root);
        while (!queue.isEmpty()) {
            List<TreeNode> nextLayer = new ArrayList<>();
            for (TreeNode node : layer) {
                node.left = create(queue.poll());
                if (node.left != null) {
                    nextLayer.add(node.left);
                }
                if (queue.isEmpty()) {
                    break;
                }
                node.right = create(queue.poll());
                if (node.right != null) {
                    nextLayer.add(node.right);
                }
            }
            layer = nextLayer;
        }
        return root;
    }

    private static TreeNode create(Integer num) {
        if (num == null) {
            return null;
        }
        return new TreeNode(num);
    }

    public void print() {
        Queue<TreeNode> queue = new LinkedList<>();
        queue.offer(this);
        boolean isAllNull = false;

        while (!queue.isEmpty()) {
            if (isAllNull) {
                break;
            }
            isAllNull = true;
            int levelSize = queue.size();
            for (int i = 0; i < levelSize; i++) {
                TreeNode node = queue.poll();
                if (node != null) {
                    System.out.printf("%-6d", node.val);
                    queue.offer(node.left);
                    queue.offer(node.right);
                    isAllNull = node.left == null && node.right == null && isAllNull;
                } else {
                    queue.offer(null);
                    System.out.print("#     ");
                }
            }
            System.out.println();
        }
    }
}
