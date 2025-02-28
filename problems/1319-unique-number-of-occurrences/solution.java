class Solution {
    public boolean uniqueOccurrences(int[] arr) {
        Set<Integer> h = new HashSet<>();
        Arrays.sort(arr);
        int num = 0;
        int counter = 0;
        for (int i = 0; i < arr.length; i++) {
            if (num != arr[i]){
                if (h.contains(counter)){
                    return false;
                }
                h.add(counter);
                num = arr[i];
                counter = 1;
            }
            else {
                counter ++;
            }
        }
        if (h.contains(counter)){
            return false;
        }
        return true;
    }
}