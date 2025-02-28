


class Solution {
    public List<Integer> findDuplicates(int[] nums) {
        List<Integer> h = new ArrayList<>();
        int g = 0;
        for (int i = 0; i < nums.length; i++) {
            g = Math.abs(nums[i]) - 1;
            if (nums[g] < 0){
                h.add(Math.abs(g) + 1);
            }
            nums[g] *= -1;
        }
        return h;
    }
}