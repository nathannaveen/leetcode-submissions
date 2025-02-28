class Solution {
    public boolean checkRecord(String s) {
        int absentCounter = 0;
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == 'A'){
                absentCounter ++;
                if (absentCounter == 2) return false;
            }
            if (i <= s.length() - 3 && s.charAt(i) == 'L' && s.charAt(i + 1) == 'L' && s.charAt(i + 2) == 'L'){
                return false;
            }
        }

        return true;
    }
}