class Solution {
    public int climbStairs(int n) {
        if (n <= 2){
            return n;
        }

        int num = 1;
        int old1 = 1;
        int old2 = 0;

        for (int i = 2; i <= n; i++) {
            old2 = old1;
            old1 = num;
            num += old2;
        }
        return num;
    }
}