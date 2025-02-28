class Solution {
    public int[][] flipAndInvertImage(int[][] A) {
        for (int i = 0; i < A.length; i++) {
            int left = 0;
            int right = A[i].length - 1;
            while (left < right){
                int g = A[i][right];
                A[i][right] = A[i][left];
                A[i][left] = g;
                left ++;
                right --;
            }
            for (int j = 0; j < A[i].length; j++) {
                switch (A[i][j]){
                    case 0:
                        A[i][j] = 1;
                        break;
                    case 1:
                        A[i][j] = 0;
                        break;
                }
            }
        }
        return A;
    }
}