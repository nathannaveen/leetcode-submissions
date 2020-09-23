class Solution {
    public int[] twoSum(int[] nums, int target) {
        HashMap<Integer, Integer> h = new HashMap<>();
        for (int j = 0; j < nums.length; j++) {
            if (h.containsKey(target - nums[j])){
                return new int[]{h.get(target - nums[j]), j};
            }
            h.put(nums[j], j);
        }
        return null;
    }
}