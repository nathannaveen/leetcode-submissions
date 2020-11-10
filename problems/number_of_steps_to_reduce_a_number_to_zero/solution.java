class Solution {
    public int numberOfSteps (int num) {
        if (num == 1){
            return 1;
        }
        if (num <= 0){
            return 0;
        }
        int c = 0;
        while (num > 2){
            if (num %2 == 0){
                num/=2;
                c++;
            }
            else{
                num -= 1;
                c++;
            }
        }
        return c+2;
    }
}