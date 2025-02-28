class Solution {
    public boolean checkPossibility(int[] nums) {
        boolean one = false;
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] < nums[i - 1]) {
                if (i != 1 && nums[i] < nums[i - 2]) nums[i] = nums[i - 1];
                if (one) return false;
                one = true;
            }
        }
        return true;
    }
}