class Solution {
    public boolean backspaceCompare(String S, String T) {
        if (g(S).equals(g(T))){
            return true;
        }
        else{
            return false;
        }

    }
    public static Stack<Character> g(String h){
        Stack<Character> x = new Stack<>();
        for (int i = 0; i < h.length(); i++) {
            if (h.charAt(i) == '#' && x.size() > 0){
                x.pop();
            }
            else if (h.charAt(i) != '#'){
                x.push(h.charAt(i));
            }
        }
        return x;
    }
}