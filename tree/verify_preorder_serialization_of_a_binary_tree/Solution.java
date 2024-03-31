package tree.verify_preorder_serialization_of_a_binary_tree;

import java.util.Stack;

/**
 * @author Bota5ky
 * @link <a href="https://leetcode.cn/problems/verify-preorder-serialization-of-a-binary-tree/">331. 验证二叉树的前序序列化</a>
 * @since 2024-03-31 12:26
 */
class Solution {
    public boolean isValidSerialization(String preorder) {
        Stack<String> stack = new Stack<>();
        String[] split = preorder.split(",");
        for (String s : split) {
            while (stack.size() >= 2 && s.equals("#") && stack.peek().equals("#")) {
                stack.pop();
                String pop = stack.pop();
                if (pop.equals("#")) {
                    return false;
                }
            }
            stack.push(s);
        }
        return stack.size() == 1 && stack.pop().equals("#");
    }
}
