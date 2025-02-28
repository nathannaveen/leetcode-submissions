class Solution {
    public int heightChecker(int[] heights) {
        int counter = 0;
        int[] g = new int[heights.length];

        for (int i = 0; i < heights.length; i++) {
            g[i] = heights[i];
        }
        
        Arrays.sort(heights);

        for (int i = 0; i < heights.length; i++) {
            if (heights[i] != g[i]){
                counter ++;
            }
        }
        return counter;
    }
}