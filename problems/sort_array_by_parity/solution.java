class Solution {
    public int[] sortArrayByParity(int[] A) {
        int[] h = new int[A.length];
        int hCounter = 0;
        List<Integer> odd = new ArrayList<>();
        for (int i = 0; i < A.length; i++) {
            if (A[i] % 2 == 0){
                h[hCounter] = A[i];
                hCounter ++;
            }
            else {
                odd.add(A[i]);
            }
        }
        for (int i = 0; i < odd.size(); i++) {
            h[hCounter] = odd.get(i);
            hCounter ++;
        }
        return h;
    }
}