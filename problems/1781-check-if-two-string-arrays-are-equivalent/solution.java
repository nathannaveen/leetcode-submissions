class Solution {
    public boolean arrayStringsAreEqual(String[] word1, String[] word2) {
        String w1 = "";
        String w2 = "";
        for (String i : word1){
            w1 += i;
        }
        for (String i : word2){
            w2 += i;
        }
        return w1.equals(w2);
    }
}