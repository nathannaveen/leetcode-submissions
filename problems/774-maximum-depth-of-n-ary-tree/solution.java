/*
// Definition for a Node.
class Node {
    public int val;
    public List<Node> children;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val, List<Node> _children) {
        val = _val;
        children = _children;
    }
};
*/

class Solution {
    public int TheMaxDepth = 0;
    public int maxDepth(Node root) {
        if (root == null) return 0;
        getMaxDepth(root, 1);
        return TheMaxDepth;
    }
    public void getMaxDepth(Node node, int depth){
        if (node == null) return;
        TheMaxDepth = Math.max(TheMaxDepth, depth);
        for (Node theChild : node.children){
            getMaxDepth(theChild, depth + 1);
        }
    }
}