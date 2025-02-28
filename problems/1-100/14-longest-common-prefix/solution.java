class Solution {
    public String longestCommonPrefix(String[] strs) {
        if(strs == null || strs.length == 0)    
            return "";
        String before = strs[0];
        int i = 1;
        while(i < strs.length){
            while(strs[i].indexOf(before) != 0)
                before = before.substring(0,before.length()-1);
            i++;
        }
        return before;
    }
}