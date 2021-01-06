class Solution {
   public int diagonalSum(int[][] mat) {
        int sum = 0;
        boolean odd = false;
        int row = 0;

        if (mat.length % 2 == 1){
            odd = true;
        }

        for (int col = 0; col < mat.length; col++) {
            sum += mat[row][col];
            row++;
        }

        row = 0;
        for (int col = mat.length - 1; col >= 0; col--) {
            sum += mat[row][col];
            row++;
        }

        if (odd){
            sum -= mat[mat.length/2][mat.length/2];
        }

        return sum;
    }
}