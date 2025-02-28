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
    public boolean isSymmetric(TreeNode root) {
        Stack<TreeNode> left = new Stack<>();
        Stack<TreeNode> right = new Stack<>();
        if (root == null){
            return true;
        }
        left.add(root.left);
        right.add(root.right);
        while (!right.isEmpty() || !left.isEmpty()){
            TreeNode l = left.pop();
            TreeNode r = right.pop();
            if (l != null && r != null) {
                if (r.val != l.val) {
                    return false;
                }
                left.add(l.left);
                left.add(l.right);
                right.add(r.right);
                right.add(r.left);
            }
            else if ((l != null && r == null) || (l == null && r != null)) {
                return false;
            }
        }
        return true;
    }
}