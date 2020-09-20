class Solution {
    public boolean isPowerOfFour(int num) {
        return Integer.toString(num, 4).matches("^10*$");
    }
}