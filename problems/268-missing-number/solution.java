class Solution {
    public int missingNumber(int[] nums) {
        Arrays.sort(nums);
        if (nums[0] != 0){
            return 0;
        }
        
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] - nums[i-1] == 2){
                return nums[i] - 1;
            }
        }
        return nums[nums.length-1] +1;
        
    }
}