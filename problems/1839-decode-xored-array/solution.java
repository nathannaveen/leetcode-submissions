class Solution {
    public int[] decode(int[] encoded, int first) {
        int[] h = new int[encoded.length + 1];
        h[0] = first;

        for (int i = 1; i < h.length; i++) {
            h[i] = encoded[i - 1] ^ h[i - 1];
        }

        return h;
    }
}