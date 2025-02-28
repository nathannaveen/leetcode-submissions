class Solution {
    public String removeDuplicates(String S) {
        boolean pops = true;

        while (pops){
            boolean p = true;
            Stack<Character> h = new Stack<>();
            for (int i = 0; i < S.length(); i++) {
                if (h.size() > 0 && h.peek() == S.charAt(i)){
                    h.pop();
                    p = false;
                }
                else {
                    h.push(S.charAt(i));

                }
            }
            if (!p){
                p = true;
            }
            else {
                pops = false;
            }
            S = "";
            for (int i = 0; i < h.size(); i++) {
                S += h.get(i);
            }
            System.out.println(S);
        }
        return S;
    }
}