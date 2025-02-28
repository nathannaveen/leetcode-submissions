/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public boolean isUnivalTree(TreeNode root) {
        Stack<TreeNode> h = new Stack<>();
        int num = root.val;
        h.push(root);
        while (h.size() != 0){
            TreeNode g = h.pop();
            if (g != null) {
                if (g.val != num) {
                    return false;
                }
                h.push(g.right);
                h.push(g.left);
            }
        }
        return true;
    }
}