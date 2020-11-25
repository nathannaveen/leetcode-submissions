class Solution {
    public int[] smallerNumbersThanCurrent(int[] nums) {
        int[] h = new int[nums.length];
        for (int i = 0; i < nums.length; i++) {
            int g = 0;
            for (int j = 0; j < nums.length; j++) {
                if (nums[i] > nums[j]) {
                    g ++;
                }
            }
            h[i] = g;
        }
        return h;
    }
}