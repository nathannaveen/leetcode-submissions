class Solution {
    public int countLargestGroup(int n) {
        int[] h = new int[n + 2];
        for (int i = 1; i <= n; i++) {
            int g = i;
            int sum = 0;
            while (g > 0){
                sum += g % 10;
                g /= 10;
            }
            h[sum] ++;
        }

        Arrays.sort(h);
        int num = h[h.length - 1];
        int counter = 0;

        for (int i = h.length - 1; i > 0; i--) {
            if (h[i] == num){
                counter ++;
            }
            else {
                return counter;
            }
        }

        return 0;
    }
}