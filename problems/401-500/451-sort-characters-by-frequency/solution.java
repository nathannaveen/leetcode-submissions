class Solution {
    public String frequencySort(String s) {
        HashMap<Character, Integer> h = new HashMap<>();
        for (char i : s.toCharArray()) {
            if (h.containsKey(i)) h.replace(i, h.get(i) + 1);
            else h.put(i, 1);
        }
        s = "";
        for (int i = 0; i < h.size(); i++) {
            int max = 0;
            char letter = 0;
            for (char j : h.keySet()){
                if (h.get(j) > max) {
                    max = h.get(j);
                    letter = j;
                }
            }
            for (int j = 0; j < h.get(letter); j++) {
                s += letter;
            }
            h.replace(letter, -1);
        }
        return s;
    }
}