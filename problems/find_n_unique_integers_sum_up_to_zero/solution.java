class Solution {
    public int[] sumZero(int n) {
        int[] h = new int[n];
        h[0] = ((n-2)*((n-2) + 1))/2;
        int counter = 0;
        for (int i = 1; i < n; i++) {
            h[i] = counter*(-1);
            counter ++;
        }
        return h;
    }
}
