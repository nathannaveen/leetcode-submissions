class Solution {
    public int countNegatives(int[][] grid) {
        int counter = 0;
        for (int row = 0; row < grid.length; row++) {
            for (int col = grid[row].length - 1; col >= 0; col--) {
                if (grid[row][col] < 0){
                    counter ++;
                }
                else {
                    break;
                }
            }
        }
        return counter;
    }
}