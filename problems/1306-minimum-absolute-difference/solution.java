class Solution {
    public List<List<Integer>> minimumAbsDifference(int[] arr) {
        List<List<Integer>> res = new ArrayList<>();
        Arrays.sort(arr);
        int min = Integer.MAX_VALUE;
        for (int i = 0; i < arr.length - 1; i++) {
            if (arr[i + 1] - arr[i] < min){
                min = arr[i + 1] - arr[i];
            }
        }
        for (int i = 0; i < arr.length - 1; i++) {
            if (arr[i + 1] - arr[i] == min){
                List<Integer> h = new ArrayList<>();
                h.add(arr[i]);
                h.add(arr[i + 1]);
                res.add(h);
            }
        }
        return res;
    }
}