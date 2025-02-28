class Solution {
    public int maxRepeating(String sequence, String word) {
        for (int i = sequence.length() / word.length(); i > 0; i--) {
            if (sequence.indexOf(word.repeat(i)) >= 0) {
                return i;
            }
        }
        return 0;
    }
}