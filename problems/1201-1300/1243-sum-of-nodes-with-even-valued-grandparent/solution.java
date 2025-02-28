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
    public int sumEvenGrandparent(TreeNode root) {
        int sum = 0;
        Stack<TreeNode> h = new Stack<>();
        h.push(root);
        while (!h.isEmpty()){
            TreeNode g = h.pop(); // grandparent
            if (g != null) { // checking whether it is even
                if (g.val % 2 == 0) {
                    TreeNode l = g.left; // sons and daughters on the left
                    TreeNode r = g.right; // sons and daughter on the right
                    if (l != null) {
                        if (l.left != null) { // l.left is the grandchildren on the left left
                            sum += l.left.val;
                        }
                        if (l.right != null) { // l.right is the grandchildren on the left right
                            sum += l.right.val;
                        }
                    }
                    if (r != null) {
                        if (r.left != null) { // r.left is the grandchildren on the right left
                            sum += r.left.val;
                        }
                        if (r.right != null) { // r.right is the grandchildren on the right right
                            sum += r.right.val;
                        }
                    }
                }
                h.push(g.right); // add the right to the stack to continue the loop
                h.push(g.left); // add the left to the stack to continue the loop
            }

        }
        return sum;
    }
}