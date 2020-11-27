import java.util.Arrays;

class Solution {
    public int arrayPairSum(int[] nums) {
        int sum = 0;
        Arrays.sort(nums);
        for (int i = 0; i < nums.length; i += 2) {
            sum += Math.min(nums[i], nums[i+1]);
        }
        return sum;
    }
}