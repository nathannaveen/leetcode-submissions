class Solution {
    public int maxDepth(String s) {
        Stack<String> h = new Stack<>();
        int max = 0;
        for (char i : s.toCharArray()){
            if (i == '('){
                h.push("(");
                if (h.size() > max){
                    max = h.size();
                }
            }
            else if (i == ')'){
                h.pop();
            }
        }
        return max;
    }
}