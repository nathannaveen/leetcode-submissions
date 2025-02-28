class Solution {
    public int sumOddLengthSubarrays(int[] arr) {
        int total = 0;
        for (int i = 1; i < arr.length; i+= 2) {
            for (int j = 0; j < arr.length; j++) {
                if (j + i > arr.length){
                    break;
                }
                for (int k = j; k < j + i; k++) {
                    total += arr[k];
                }
            }
        }
        int all = 0;
        if (arr.length % 2 == 1) {
            for (int i = 0; i < arr.length; i++) {
                all += arr[i];
            }
        }
        total += all;
        return total;
    }
}