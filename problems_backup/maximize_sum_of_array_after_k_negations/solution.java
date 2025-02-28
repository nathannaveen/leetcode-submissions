class Solution {
    public int largestSumAfterKNegations(int[] A, int K) {
        int sum = 0;
        Arrays.sort(A);

        for (int i = 0; i < K; i++) {
            A[0] *= -1;
            if (A[0] > A[1]) Arrays.sort(A);
        }

        for (int i = 0; i < A.length; i++) {
            sum += A[i];
        }
        return sum;
    }
}