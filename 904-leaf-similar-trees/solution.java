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
    public boolean leafSimilar(TreeNode root1, TreeNode root2) {
        Stack<TreeNode> h = new Stack<>();
        List<Integer> leafs = new ArrayList<>();
        List<Integer> leafs2 = new ArrayList<>();
        h.push(root1);
        while (!h.isEmpty()){
            TreeNode g = h.pop();
            if (g != null){
                if (g.left == null && g.right == null){
                    leafs.add(g.val);
                }
                else {
                    h.push(g.right);
                    h.push(g.left);
                }
            }
        }
        h.push(root2);
        while (!h.isEmpty()){
            TreeNode g = h.pop();
            if (g != null){
                if (g.left == null && g.right == null){
                    leafs2.add(g.val);
                }
                else {
                    h.push(g.right);
                    h.push(g.left);
                }
            }
        }
        if (leafs.size() == leafs2.size()){
            for (int i = 0; i < leafs.size(); i++) {
                if (leafs.get(i) != leafs2.get(i)){
                    return false;
                }
            }
            return true;
        }
        return false;
    }
}