class Solution {
    public List<Integer> luckyNumbers (int[][] matrix) {
        List<Integer> h = new ArrayList<>();
        int minIndex = 0;
        
        for (int i = 0; i < matrix.length; i++) {
            int min = Integer.MAX_VALUE;
            int max = 0;
            for (int j = 0; j < matrix[i].length; j++) {
                if (matrix[i][j] < min){
                    min = matrix[i][j];
                    minIndex = j;
                }
            }
            for (int j = 0; j < matrix.length; j++) {
                if (matrix[j][minIndex] > max){
                    max = matrix[j][minIndex];
                }
            }
            if (max == min){
                h.add(max);
            }
        }
        return h;
    }
}