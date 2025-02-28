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
    public TreeNode bstToGst(TreeNode root) {
        List<Integer> total = new ArrayList<>();
        int totalSum = 0;
        Stack<TreeNode> h = new Stack<>();
//        Stack<TreeNode> j = new Stack<>();
//        j.add(root);
        h.add(root);

        while (!h.isEmpty()){
            TreeNode g = h.pop();
            if (g != null){
                total.add(g.val);
                totalSum += g.val;
                h.push(g.right);
                h.push(g.left);
            }
        }

        Collections.sort(total);
        h.push(root);

        while (!h.isEmpty()){
            TreeNode g = h.pop();
            if (g != null){
                int sum = 0;
                for (int i = 0; i < total.size(); i++) {
                    if (total.get(i) == g.val) break;
                    else sum += total.get(i);
                }
                g.val = totalSum - sum;
                h.push(g.right);
                h.push(g.left);
            }
        }

        return root;
    }
}