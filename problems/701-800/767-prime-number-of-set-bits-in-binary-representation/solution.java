class Solution {
    public boolean isPrime(int N) {
        int g = (int) Math.sqrt(N);
        if (g * g == N){
            return false;
        }

        for (int i = 2; i < Math.sqrt(N); i++) {
            if (N % i == 0){
                return false;
            }
        }

        return true;
    }

    public int countPrimeSetBits(int L, int R) {
        int res = 0;
        for (int i = L; i <= R; i++) {
            int counter = 0;
            for (int n = i; n > 0; n >>= 1) {
                counter += n & 1;
            }
            
            if (isPrime(counter)){
                res ++;
            }
        }
        return res;
    }
}