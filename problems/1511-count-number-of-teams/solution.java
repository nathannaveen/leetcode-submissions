class Solution {
    public int numTeams(int[] rating) {
        int teams = 0;
        int i = 0;
        int j = 1;
        int k = 2;

        while (i < rating.length - 2){
            if ((rating[i] < rating[j] && rating[j] < rating[k]) ||
                    (rating[i] > rating[j] && rating[j] > rating[k])){
                teams ++;
            }

            if ((j == rating.length - 2) && (k == rating.length - 1)){
                i ++;
                j = i + 1;
                k = i + 2;
            }
            else if ((k == rating.length - 1)){
                j ++;
                k = j + 1;
            }
            else {
                k ++;
            }
        }
        return teams;
    }
}