class Solution {
    public int countConsistentStrings(String allowed, String[] words) {
        int num = 0;
        Set<Character> h = new HashSet<>();
        for (char i : allowed.toCharArray()){
            h.add(i);
        }
        for (int i = 0; i < words.length; i++) {
            boolean wholeWord = true;
            for (char j : words[i].toCharArray()){
                if (!h.contains(j)){
                    wholeWord = false;
                    break;
                }
            }
            if (wholeWord){
                num += 1;
            }
        }
        return num;
    }
}