class Solution {
    public int strStr(String haystack, String needle) {
        String[] h = haystack.split("");
        String[] n = needle.split("");
        boolean g =false;
        if (needle.equals("")){
            return 0;
        }
        if (haystack.equals("")){
            return -1;
        }
        for (int i = 0; i < h.length+1 - n.length; i++) {
            if (haystack.substring(i, i + n.length).equals(needle)){
                return i;
            }
        }

        return -1;

    }
}