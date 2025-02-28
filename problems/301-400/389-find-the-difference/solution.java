class Solution {
    public char findTheDifference(final String s, final String t) {
        char result = 0;
        for (final char c : s.toCharArray()) {
            result ^= c;
        }
        for (final char c : t.toCharArray()) {
            result ^= c;
        }
        return result;
    }
}