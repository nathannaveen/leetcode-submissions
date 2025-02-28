class Solution {
    public int climbStairs(int n) {
        int[] h = new int[n + 1];
        h[0] = 1;
        h[1] = 1;
        for (int i = 2; i <= n; i++) {
            h[i] = h[i - 1] + h[i - 2];
        }
        return h[n];
    }
}