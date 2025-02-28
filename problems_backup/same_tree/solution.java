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
    public boolean isSameTree(TreeNode p, TreeNode q) {
        Stack<TreeNode> ps = new Stack<>();
        Stack<TreeNode> qs = new Stack<>();
        ps.push(p);
        qs.push(q);

        while (!ps.isEmpty()){
            TreeNode g = ps.pop();
            TreeNode m = qs.pop();

            if (g != null && m != null){
                if (g.val != m.val){
                    return false;
                }
                ps.push(g.left);
                ps.push(g.right);
                qs.push(m.left);
                qs.push(m.right);
            }
            else if (g != m){
                return false;
            }
        }
        return true;
    }
}