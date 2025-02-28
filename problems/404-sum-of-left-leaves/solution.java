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
    public int sumOfLeftLeaves(TreeNode root) {
        int sum = 0;
        Stack<TreeNode> h = new Stack<>();
        h.add(root);
        while (!h.isEmpty()){
            TreeNode g = h.pop();
            if (g != null){
                if (g.left != null && g.left.left == null && g.left.right == null)
                    sum += g.left.val;
                h.add(g.left);
                h.add(g.right);
            }
        }
        return sum;
    }
}