class Solution {
    public boolean isRectangleOverlap(int[] rec1, int[] rec2) {
        int t = rec1[0]; // I call it TACO, and TACO
        int t2 = rec2[2]; //                    ||||
        int a = rec2[0]; //                     2222
        int a2 = rec1[2];
        int c = rec1[1];
        int c2 = rec2[3];
        int o = rec2[1];
        int o2 = rec1[3];
        return t < t2 && a < a2 && c < c2 && o < o2;
        

    }
}