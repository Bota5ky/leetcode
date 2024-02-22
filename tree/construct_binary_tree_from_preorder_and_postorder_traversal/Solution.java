package tree.construct_binary_tree_from_preorder_and_postorder_traversal;

import _model.TreeNode;

import java.util.Arrays;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-postorder-traversal/">889. 根据前序和后序遍历构造二叉树</a>
 * @since 2024-02-22 08:17
 */
class Solution {
    public TreeNode constructFromPrePost(int[] preorder, int[] postorder) {
        if (preorder.length == 0) {
            return null;
        }
        int idx = 0;
        TreeNode root = new TreeNode(preorder[0]);
        if (preorder.length == 1) {
            return root;
        }
        while (postorder[idx] != preorder[1]) {
            idx++;
        }
        root.left = constructFromPrePost(Arrays.copyOfRange(preorder, 1, idx + 2),
                Arrays.copyOfRange(postorder, 0, idx + 1));
        root.right = constructFromPrePost(Arrays.copyOfRange(preorder, idx + 2, preorder.length),
                Arrays.copyOfRange(postorder, idx + 1, postorder.length - 1));
        return root;
    }
}
