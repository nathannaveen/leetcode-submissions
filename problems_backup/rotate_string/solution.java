class Solution {
    public boolean rotateString(String A, String B) {
        if (A.equals(B)){
            return true;
        }
        
        char[] h = A.toCharArray();
        char[] l = B.toCharArray();
        
        for (int i = 0; i < A.length(); i++) {
            char m = h[0];
            for (int j = 1; j < A.length(); j++) {
                h[j - 1] = h[j];
            }
            h[A.length() - 1] = m;
            if (Arrays.equals(h, l)){
                return true;
            }
        }
        
        return false;
    }
}