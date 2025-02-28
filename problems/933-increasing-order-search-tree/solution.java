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
    public TreeNode increasingBST(TreeNode root) {
        Stack<TreeNode> k = new Stack<>();
        List<Integer> h = new ArrayList<>();
        k.add(root);
        while (!k.isEmpty()){
            TreeNode g = k.pop();
            if (g != null) {
                h.add(g.val);
                k.add(g.right);
                k.add(g.left);
            }
        }
        Collections.sort(h);
        TreeNode n = new TreeNode(h.get(0));
        k.add(n);
        for (int i = 1; i < h.size(); i++) {
            TreeNode g = k.pop();
            if (g != null){
                g.right = new TreeNode(h.get(i));
                k.add(g.right);
            }
        }
        return n;
    }
}