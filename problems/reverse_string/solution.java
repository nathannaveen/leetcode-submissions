class Solution {
    public void reverseString(char[] s) {
        for (int i = 0; i < s.length / 2; i++) {
            char oldsOfI = s[i];
            s[i] = s[s.length-1-i];
            s[s.length-1-i] = oldsOfI;

        }
        System.out.println(Arrays.toString(s));
    }
}