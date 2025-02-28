class Solution {
    public int removeElement(int[] nums, int val) {
        int counter = 0;
        for (int i : nums){
            if (i != val){
                nums[counter] = i;
                counter ++;
            }
        }
        return counter;
    }
}
