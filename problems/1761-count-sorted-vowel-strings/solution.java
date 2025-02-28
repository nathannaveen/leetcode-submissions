class Solution {
    public int countVowelStrings(int n) {
        int[] h = new int[5];
        int sum = 0;

        h[0] = 1; h[1] = 1; h[2] = 1; h[3] = 1; h[4] = 1;

        for (int i = 0; i < n-1; i++) {
            for (int j = 1; j < 5; j++) {
                h[j] += h[j - 1];
            }
        }

        for (int i = 0; i < h.length; i++) {
            sum += h[i];
        }
        return sum;
    }
}