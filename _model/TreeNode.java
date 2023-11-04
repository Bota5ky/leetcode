package _model;

import java.util.ArrayList;

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

    public static TreeNode create(int[] nums) {
        if (nums.length == 0) {
            return null;
        }
        TreeNode root = new TreeNode(nums[0]);
        ArrayList<TreeNode> layer = new ArrayList<TreeNode>();
        layer.add(root);
        for (int i = 1; i < nums.length;) {
            ArrayList<TreeNode> nextLayer = new ArrayList<TreeNode>();
            for (TreeNode treeNode : layer) {
                treeNode.left = new TreeNode(nums[i++]);
                nextLayer.add(treeNode.left);
                if (i < nums.length) {
                    treeNode.right = new TreeNode(nums[i++]);
                    nextLayer.add(treeNode.right);
                }
            }
            layer = nextLayer;
        }
        return root;
    }
}
