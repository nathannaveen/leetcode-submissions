class Solution {
    public int[] decrypt(int[] code, int k) {
        int[] h = new int[code.length];
        if (k > 0) {
            for (int i = 0; i < code.length; i++) {
                int g = i;
                int sum = 0;
                for (int j = 1; j <= k; j++) {
                    if (i + j == code.length) {
                        g = 0;
                    } else {
                        g++;
                    }
                    sum += code[g];
                }
                h[i] = sum;
            }
        }
        else if (k == 0){
            for (int i = 0; i < code.length; i++) {
                h[i] = 0;
            }
        }
        else {
            for (int i = 0; i < code.length; i++) {
                int g = i;
                int sum = 0;
                for (int j = 1; j <= (-1) * k; j++) {
                    if (i - j < 0 && g == 0){
                        g = code.length - 1;
                    }
                    else {
                        g --;
                    }
                    sum += code[g];
                }
                h[i] = sum;
            }
        }
        return h;
    }
}