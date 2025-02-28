class Solution {
    public int findNumbers(int[] nums) {
        int totalCounter = 0;
        for (int i = 0; i < nums.length; i++) {
            int counter = 0;
            int g = nums[i];
            while (g > 0){
                g /= 10;
                counter ++;
            }
            if (counter % 2 == 0){
                totalCounter += 1;
            }
        }
        return totalCounter;
    }
}