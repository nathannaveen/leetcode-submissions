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
    public int findSecondMinimumValue(TreeNode root) {
        Stack<TreeNode> h = new Stack<>();
        List<Integer> l = new ArrayList<>();
        h.add(root);

        while (!h.isEmpty()){
            TreeNode g = h.pop();

            if (g != null){
                l.add(g.val);
                h.add(g.left);
                h.add(g.right);
            }
        }

        Collections.sort(l);
        
        int num = l.get(0);
        for (int i = 1; i < l.size(); i++) {
            if (l.get(i) != num){
                return l.get(i);
            }
        }
        return -1;
    }
}