class Solution {
    public int numJewelsInStones(String J, String S) {
        int counter = 0;
        for (int i = 0; i < S.length(); i++) {
            for (int j = 0; j < J.length(); j++) {
                if (S.charAt(i) == J.charAt(j)){
                    counter ++;
                    break;
                }
            }
        }
        return counter;
    }
}