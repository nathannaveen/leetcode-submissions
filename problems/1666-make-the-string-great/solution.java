class Solution {
    public String makeGood(String s) {
        String result = "";
        
        Stack<Character> h = new Stack<>(); 
        boolean g = false; 
        boolean noPop = true;
        for (int i = 0; i < s.length(); i++) {
            if (h.size() > 0){ 
                if (((int) s.charAt(i) -32 == (int) h.peek()) || ((int) s.charAt(i) + 32 == (int) h.peek())){
                    h.pop();
                    noPop = false;
                }  
                else{
                    g = true;
                }
            }
            else{
                g = true;
            }
            if ((noPop == true && h.size() == 0 )|| g == true){
                h.add(s.charAt(i));
                g = false;
            }
        }
        for (int i = 0; i < h.size(); i++) {
            result += h.get(i);
        }
        return result;
    }
}