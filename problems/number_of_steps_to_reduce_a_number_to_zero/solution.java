class Solution {
    public int numberOfSteps (int num) {
        int counter = 0;
        while (num > 0){

            if ((num & 1)== 0) {
                num >>= 1;
            }
            else {
                num -= 1;
            }
            counter ++;
        }
        return counter;
    }
}
