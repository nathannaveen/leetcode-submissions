class Solution {
    public int findLengthOfLCIS(int[] nums) {
        int maxCounter = 0;
        int counter = 1;
        
        if (nums.length == 0){
            return 0;
        }
        
        for (int i = 1; i < nums.length; i++) {
            if (nums[i - 1] >= nums[i]){
                maxCounter = Math.max(counter, maxCounter);
                counter = 1;
            }
            else counter ++;
        }
        maxCounter = Math.max(counter, maxCounter);
        return maxCounter;
    }
}