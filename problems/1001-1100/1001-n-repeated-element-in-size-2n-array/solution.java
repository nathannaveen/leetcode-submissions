class Solution {
    public int repeatedNTimes(int[] A) {
        int n = A.length/2;
        HashMap<Integer, Integer> h = new HashMap<>();
        for (int i : A) {
            
            if (h.containsKey(i)){
                int g = h.get(i);
                h.replace(i, g+1);
            }
            else{
                h.put(i, 1);
            }
        }
        for (int i : A) {
            if (h.get(i) == n){
                return i;
            }
        }

        return -1;
    }
}