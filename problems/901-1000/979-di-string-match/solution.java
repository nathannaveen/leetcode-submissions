class Solution {
    public int[] diStringMatch(String S) {
        int left = 0;
        int right = S.length();
        int[] h = new int[S.length() + 1];
        for (int i = 0; i < S.length(); i++) {
            if (S.charAt(i) == 'I'){
                h[i] = left;
                left ++;
            }
            else if (S.charAt(i) == 'D'){
                h[i] = right;
                right --;
            }
        }
        h[h.length - 1] = left;
        return h;
    }
}