class Solution {
    public int singleNumber(int[] nums) {
        HashMap<Integer, Integer> h = new HashMap<>();
        for (int i : nums) {
            if (!h.containsKey(i)){
                h.put(i, 1);
            }
            else{
                h.remove(i);
            }
        }
        for (int i : h.keySet()) {
            return i;
        }
        return -1;
    }
}