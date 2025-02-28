class Solution {
    public int[] shortestToChar(String S, char C) {
        int[] h = new int[S.length()];
        for (int i = 0; i < S.length(); i++) {
            int first = Integer.MAX_VALUE;
            int second = Integer.MAX_VALUE;
            int firstCounter = 0;
            int secondCounter = 0;
            for (int j = i; j < S.length(); j++) {
                firstCounter ++;
                if (S.charAt(j) == C){
                    first = firstCounter;
                    break;
                }
                
            }
            for (int j = i; j >= 0; j--) {
                secondCounter ++;
                if (S.charAt(j) == C){
                    second = secondCounter;
                    break;
                }
            }
            
            h[i] = Math.min(first, second) - 1;
        }
        return h;
    }
}