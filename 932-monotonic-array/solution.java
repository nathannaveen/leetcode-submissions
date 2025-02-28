class Solution {
    public boolean isMonotonic(int[] A) {
        int g = 1;
        int h = 1;
        int len = A.length;
        for (int i = 1; i < len; i++) {
            if (A[i] - A[i - 1] > 0)
                g++;
            else if (A[i] - A[i - 1] < 0)
                h++;
            else {
                g++;
                h++;
            }

        }
        return g == len || h == len;
    }
}