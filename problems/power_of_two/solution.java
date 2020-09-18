class Solution {
    public boolean isPowerOfTwo(int n) {
        if (n == 1){
            return true;
        }
        if (n == 0 || n < 0){
            return false;
        }
        int newN = n;
        while (newN > 1){
            if (newN %2  == 1){
                return false;
            }
            newN /= 2;
            
        }
        return true;
    }
}