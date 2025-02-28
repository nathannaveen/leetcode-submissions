class Solution {
    public int firstUniqChar(String s) {
        HashMap<Character, Integer> h = new HashMap<>();
        Set<Character> not = new HashSet<>();
        for (int i = 0; i < s.length(); i++) {
            if (h.containsKey(s.charAt(i))){
                h.remove(s.charAt(i));
                not.add(s.charAt(i));
            }
            else if (not.contains(s.charAt(i))){
                continue;
            }
            else {
                h.put(s.charAt(i), i);
            }
        }
        System.out.println(h.toString());
        int min = Integer.MAX_VALUE;
        for (char i : h.keySet()){
            if (h.get(i) < min){
                min = h.get(i);
            }
        }
        for (char i : h.keySet()){
            if (h.get(i) == min){
                return h.get(i);
            }
        }
        return -1;
    }
}