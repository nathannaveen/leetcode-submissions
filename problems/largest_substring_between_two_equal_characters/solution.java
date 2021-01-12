class Solution {
    public int maxLengthBetweenEqualCharacters(String s) {
        HashMap<Character, Integer> h = new HashMap<>();
        int maxCounter = 0;
        boolean contains = false;
        for (int i = 0; i < s.length(); i++) {
            if (h.containsKey(s.charAt(i))){
                contains = true;
                maxCounter = Math.max(maxCounter, (i - 1) - h.get(s.charAt(i)));
            }
            else {
                h.put(s.charAt(i), i);
            }
        }
        if (contains){
            return maxCounter;
        }
        return -1;
    }
}