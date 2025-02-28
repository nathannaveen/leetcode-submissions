class Solution {
    public int majorityElement(int[] nums) {
        Arrays.sort(nums);
        int n = nums.length;
        int num = 0;
        int counter = 0;
        for (int i = 0; i < nums.length; i++) {
            if (num != nums[i]){
                if (counter > n/2){
                    return num;
                }
                num = nums[i];
                counter = 1;
            }
            else {
                counter ++;
            }
        }
        return num;
    }
}