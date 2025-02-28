class Solution {
    public String reverseWords(String s2) {
        String ns = "";
        String s = extracted(s2);

        System.out.println(s);
        int end = s.length();
        for (int i = s.length()-1; i >= 0; i--) {

            if (s.charAt(i) == ' '){

                for (int j = i+1; j < end; j++) {
                    ns += s.charAt(j);
                }

                end = i;
                
                ns += " ";
            }
        }
        for (int i = 0; i < end; i++) {
            ns += s.charAt(i);
        }

        
        return ns;
    }
    private String extracted(String s2) {
        int i2 = s2.length() -1;
        String s = "";
        for (int i = s2.length()-1; i >=0 ; i--) {
            s += s2.charAt(i);
        }
        return s;
    }
}