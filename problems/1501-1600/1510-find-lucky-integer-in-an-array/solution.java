class Solution {
    public int findLucky(int[] arr) {
        int h = 0;
        Arrays.sort(arr);
        int num = 0;
        int counter = 0;

        for (int i = 0; i < arr.length; i++) {
            if (num != arr[i]){
                if (num == counter){
                    if (num > h){
                        h = counter;
                    }
                }
                num = arr[i];
                counter = 1;
            }
            else {
                counter ++;
            }
        }
        if (h == 0){
            h = -1;
        }
        if (num == counter){
            if (num > h){
                h = counter;
            }
        }
        return h;
    }
}