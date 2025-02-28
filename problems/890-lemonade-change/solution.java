class Solution {
    public boolean lemonadeChange(int[] bills) {
        int five = 0;
        int ten = 0;
        for (int i = 0; i < bills.length; i++) {
            if (bills[i] == 5) {
                five++;
            } else if (bills[i] == 10) {
                five--;
                ten++;
            } else if (ten > 0 && five >= 1) {
                ten--;
                five--;
            } else {
                five -= 3;
            }
            if (five < 0) {
                return false;
            }

        }
        return true;
    }
}