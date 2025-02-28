class Solution {
    public int findSpecialInteger(int[] arr) {
        int counter = 0;
        int num = 0;

        for (int i : arr){
            if (i != num){
                if (counter > arr.length/4) return num;
                num = i;
                counter = 1;
            }
            else counter ++;
        }
        return num;
    }
}