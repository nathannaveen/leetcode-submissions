class Solution {
    public boolean isPalindrome(int x) {
        int sum = 0;
        int sum2 = 0;
        int newx = x;
        int counter = 0;
        if (x < 0){
            return false;
        }
        while (newx>0){
            newx /= 10;
            counter ++;
        }
        int Newx2 = x;
        int counter2 = counter;
        for (int i = 0; i < counter2 / 2; i++) {
            int theNumber = Newx2 % 10;
            int product = theNumber;
            for (int j = 0; j < counter / 2 -1 ; j++) {
                product*=10;
            }
            sum += product;
            Newx2 /= 10;
            counter -= 2;
        }
        int Newx3 = x;
        if (counter2 % 2 == 0) {
            for (int i = 0; i < counter2 / 2; i++) {
                Newx3 /= 10;
            }
        }
        else {
            for (int i = 0; i < counter2 / 2 + 1; i++) {
                Newx3 /= 10;
            }
        }
        if (Newx3 == sum){
            return true;
        }
        else {
            return false;
        }
    }
}