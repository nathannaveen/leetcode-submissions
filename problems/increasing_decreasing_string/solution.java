class Solution {
    public String sortString(String s) {
        int[] h = new int[26];
        String n = "";

        for (char m : s.toCharArray()){
            h[((int) m) - 'a'] ++;
        }

        while (n.length() != s.length()){
            for (int j = 0; j < h.length; j++) {
                if (h[j] != 0){
                    n += ((char) (j + 'a'));
                    h[j] --;
                }
            }
            for (int j = h.length - 1; j >= 0; j--) {
                if (h[j] != 0){
                    n += ((char) (j + 'a'));
                    h[j] --;
                }
            }
        }

        return n;
    }
}