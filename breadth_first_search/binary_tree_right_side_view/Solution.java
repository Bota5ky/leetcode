package breadth_first_search.binary_tree_right_side_view;

import _model.TreeNode;

import java.util.ArrayList;
import java.util.List;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/WNC0Lk/">LCR 046. 二叉树的右视图</a>
 * @since 2023-10-27 14:35
 */
class Solution {
    public List<Integer> rightSideView(TreeNode root) {
        ArrayList<TreeNode> list = new ArrayList<TreeNode>();
        if (root != null) {
            list.add(root);
        }
        ArrayList<Integer> res = new ArrayList<Integer>();
        while (!list.isEmpty()) {
            res.add(list.get(0).val);
            ArrayList<TreeNode> temp = new ArrayList<TreeNode>();
            for (TreeNode node : list) {
                TreeNode right = node.right;
                if (right != null) {
                    temp.add(right);
                }
                TreeNode left = node.left;
                if (left != null) {
                    temp.add(left);
                }
            }
            list = temp;
        }
        return res;
    }
}
