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
    public List<String> binaryTreePaths(TreeNode root) {
        List<String> res = new ArrayList<>();
        
        if (root == null) return res;
        
        String path = Integer.toString(root.val);
        if (root.right == null && root.left == null) res.add(path);
        if (root.right != null) dfs(root.right, path, res);
        if (root.left != null) dfs(root.left, path, res);
        return res;
    }
    
    public void dfs(TreeNode h, String path, List<String> res){
        path += "->" + h.val;
        
        if (h.left == null && h.right == null){
            res.add(path);
            return;
        }

        if (h.right != null) dfs(h.right, path, res);
        if (h.left != null) dfs(h.left, path, res); 
    }
}