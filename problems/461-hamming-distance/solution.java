class Solution {
    public int hammingDistance(int x, int y) {
        int num = 0;
        while (Math.max(x, y) > 0){
            
            if (x % 2 != y % 2){
                num += 1;
            }
            x >>= 1;
            y >>= 1;
        }
        return num;
    }
}