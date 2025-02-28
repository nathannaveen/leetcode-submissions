class Solution {
    public int numWaterBottles(int numBottles, int numExchange) {
        int result = numBottles;
        while (numBottles >= numExchange) {
            int mod = numBottles % numExchange;
            numBottles /= numExchange;
            result += numBottles;
            numBottles += mod;
        }
        return result;
    }
}