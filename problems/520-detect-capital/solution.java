class Solution {
    public boolean detectCapitalUse(String word) {
        for (int i = 1; i < word.length(); i++) {
            if ((int) word.charAt(i) <= 90 && (int) word.charAt(i) >= 65){
                if ((int) word.charAt(i - 1) > 90 || (int) word.charAt(i - 1) < 65){
                    return false;
                }
            }
            else if (i >= 2){
                if ((int) word.charAt(i - 1) <= 90 && (int) word.charAt(i - 1) >= 65){
                    return false;
                }
            }
        }
        return true;
    }
}