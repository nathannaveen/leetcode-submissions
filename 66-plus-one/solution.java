class Solution {
    public int[] plusOne(int[] digits) {
        int c = 0;
        int n = digits.length;

        for (int i = n - 1; i >= 0; i--) {
            int num;
            if (i == n - 1) {

                num = (1 + c + digits[i]) % 10;
                c = (1 + c + digits[i]) / 10;
            }
            else {
                num = (c + digits[i]) % 10;
                c = (c + digits[i]) / 10;
            }
            digits[i] = num;
        }

        if (c != 0) {
            int[] list = new int[n + 1];
            list[0] = c;
            System.arraycopy(digits, 0, list, 1, n);
            return list;
        }
        return digits;
    }
}