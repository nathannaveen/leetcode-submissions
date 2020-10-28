class Solution {
    public static int mySqrt(int x) {
        int theNew = 0;
        int theOld = 1;
        float oldValue = 0;
        float newValue = 10;
        if (x == 0){
            return 0;
        }
        for (int i = 2; i < x; i++) {
            theNew = i;
            newValue = x/i;
            if (oldValue > theOld && newValue < theNew){
                return theOld;
            }
            if (newValue == theNew){
                return theNew;
            }
            oldValue = newValue;
            theOld = theNew;
        }
        return 1;
    }
}