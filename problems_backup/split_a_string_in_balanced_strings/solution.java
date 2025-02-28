class Solution {
    public int balancedStringSplit(String s) {
        int counter = 0;
        int l = 0, r = 0;

        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == 'R') r ++;
            else l ++;
            if (l == r){
                counter++;
                l = 0;
                r = 0;
            }
        }
        return counter;
    }
}