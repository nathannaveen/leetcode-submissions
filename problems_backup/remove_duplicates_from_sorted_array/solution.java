class Solution {
    public int removeDuplicates(int[] nums) {
        int n = 0;
        int prev = Integer.MAX_VALUE;
        for (int i : nums){
            if (i != prev){
                nums[n++] = i;
                prev = i;
            }
        }
        return n;
    }
}