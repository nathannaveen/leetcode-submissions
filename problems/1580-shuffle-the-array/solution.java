class Solution {
    public int[] shuffle(int[] nums, int n) {
        int[] x = new int[2*n];
        int c = 0;
        for (int i = 0; i < 2*n; i+=2) {
            x[i] = nums[c];
            x[i+1] = nums[n+c];
            c++;
        }
        return x;
    }
}