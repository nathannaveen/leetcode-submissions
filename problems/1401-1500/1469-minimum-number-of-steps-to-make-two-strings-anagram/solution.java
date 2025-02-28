class Solution {
    public int minSteps(String s, String t) {
        int[] first = new int[26];
        int[] second = new int[26];
        int sum = 0;
        for (int i = 0; i < s.length(); i++) {
            first[s.charAt(i) - 97] += 1;
            second[t.charAt(i) - 97] += 1;
        }

        for (int i = 0; i < 26; i++) {
            sum += Math.abs(first[i] - second[i]);
        }
        return sum / 2;
    }
}