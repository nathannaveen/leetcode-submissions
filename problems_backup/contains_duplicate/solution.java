class Solution {
    public boolean containsDuplicate(int[] nums) {
        HashMap<Integer, Integer> h = new HashMap<>();
        for (int i : nums) {
            if (h.containsKey(i)){
                return true;
            }
            h.put(i, 1);
        }
        return false;
    }
}