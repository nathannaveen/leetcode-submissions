class Solution {
    public int maximumWealth(int[][] accounts) {
        int greatest = 0;
        for (int i = 0; i < accounts.length; i++) {
            int sum = 0;
            for (int j : accounts[i]){
                sum += j;
            }
            if (sum > greatest){
                greatest = sum;
            }
        }
        return greatest;
    }
}