class Solution {
    public boolean validMountainArray(int[] arr) {
        boolean peak = false;
        if (arr.length < 3){
            return false;
        }
        for (int i = 1; i < arr.length; i++) {
            if (!peak && arr[i] <= arr[i - 1]) return false;
            
            if (i < arr.length - 1 && !peak && arr[i] > arr[i - 1] && arr[i] > arr[i + 1]) peak = true;

            else if (peak && arr[i] >= arr[i - 1]) return false;
        }
        if (!peak) return false;
        return true;
    }
}