class Solution {
    public int fib(int N) {
        int fibinachi = 0;
        int fib2 = 1;
        int fib3 = 0;
        for (int i = 0; i < N; i++) {
            fib3 = fibinachi + fib2;
            fibinachi = fib2;
            fib2 = fib3;
        }
        return fibinachi;
    }

}