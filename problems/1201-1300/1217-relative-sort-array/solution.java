class Solution {
    public int[] relativeSortArray(int[] arr1, int[] arr2) {
        HashMap<Integer, Integer> h = new HashMap<>();
        List<Integer> g = new ArrayList<>();
        int arrOneCounter = 0;

        for (int n : arr2) {
            h.put(n, 0);
        }

        for (int n : arr1){
            if (h.containsKey(n)) h.replace(n, h.get(n) + 1);
            else g.add(n);
        }

        for (int n : arr2){
            for (int i = 0; i < h.get(n); i++) {
                arr1[arrOneCounter] = n;
                arrOneCounter ++;
            }
        }

        Collections.sort(g);
        
        for (int i = 0; i < g.size(); i++) {
            arr1[arrOneCounter] = g.get(i);
            arrOneCounter ++;
        }
        return arr1;
    }
}