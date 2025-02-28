class Solution {
    public int[] sortArrayByParityII(int[] A) {
        List<Integer> odd = new ArrayList<>();
        List<Integer> even = new ArrayList<>();
        boolean parity = true;
        for (int i = 0; i < A.length; i++) {
            if (A[i] % 2 == 0){
                even.add(A[i]);
            }
            else {
                odd.add(A[i]);
            }
        }
        int oddCounter = 0;
        int evenCounter = 0;
        for (int i = 0; i < A.length; i++) {
            if (parity){
                A[i] = even.get(evenCounter);
                evenCounter ++;
                parity = false;
            }
            else {
                A[i] = odd.get(oddCounter);
                oddCounter ++;
                parity = true;
            }
        }
        return A;
    }
}