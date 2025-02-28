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
    public boolean findTarget(TreeNode root, int k) {
        List<Integer> h = new ArrayList<>();
        Stack<TreeNode> m = new Stack<>();
        m.add(root);
        while (!m.isEmpty()){
            TreeNode g = m.pop();
            if (g != null) {
                h.add(g.val);
                m.add(g.left);
                m.add(g.right);
            }
        }
        Collections.sort(h);
        int right = h.size() - 1;
        int left = 0;
        while (left < right){
            if (h.get(left) + h.get(right) > k){
                right --;
            }
            else if (h.get(left) + h.get(right) < k){
                left ++;
            }
            else {
                return true;
            }
        }
        return false;
    }
}