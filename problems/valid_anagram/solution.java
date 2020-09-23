class Solution {
    public boolean isAnagram(String s, String t) {
        if (s.length() != t.length()){
            return false;
        }
        char[] first = s.toCharArray();
        char[] second = t.toCharArray();
        int[] letters = new int[26];
        for (int i = 0; i < first.length; i++) {
            letters[first[i] - 'a'] += 1;
            letters[second[i] - 'a'] -= 1;
        }
        for (int i : letters) {
            if (i != 0){
                return false;
            }
        }
        return true;
    }
}