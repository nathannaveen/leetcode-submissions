class Solution {
    public int addDigits(int num) {
        int g= num;
        int sum = 0;
        while (g>0){
            sum += g%10;
            g /= 10;
        }

        if (sum > 9){
            sum = addDigits(sum);
        }

        return sum;
    }
}