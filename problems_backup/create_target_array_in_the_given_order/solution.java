class Solution {
    public int[] createTargetArray(int[] nums, int[] index) {
        int[] h = new int[nums.length];
        for (int i = 0; i < index.length; i++) {
            if (h[index[i]] != 0){
                int m = h.length - 2;
                for (int j = h.length - 1; j > index[i]; j--) {
                    h[m + 1]= h[m];
                    m -= 1;
                }
                h[index[i]] = nums[i];
            }
            else {
                h[index[i]] = nums[i];
            }
        }
        return h;
    }
}