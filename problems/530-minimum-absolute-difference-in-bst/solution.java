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
    public int getMinimumDifference(TreeNode root) {
        Stack<TreeNode> h = new Stack<>();
        List<Integer> m = new ArrayList<>();
        h.add(root);
        while (!h.isEmpty()){
            TreeNode l = h.pop();
            if (l != null) {
                m.add(l.val);
                h.add(l.left);
                h.add(l.right);
            }
        }
        Collections.sort(m);
        int min = Integer.MAX_VALUE;
        for (int i = 0; i < m.size() - 1; i++) {
            int difference = Math.abs(m.get(i) - m.get(i + 1));
            if (difference < min){
                min = difference;
            }
        }
        return min;
    }
}