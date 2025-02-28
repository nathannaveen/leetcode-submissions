class Solution {
    public int minAddToMakeValid(String S) {
        Stack<Character> h = new Stack<>();
        for (char i : S.toCharArray()){
            if (h.size() > 0 && h.get(h.size() - 1) == '(' && i == ')'){
                System.out.println("pop");
                h.pop();
            }
            else {
                System.out.println("push");
                h.push(i);
            }
        }
        return h.size();
    }
}