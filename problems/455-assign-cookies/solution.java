class Solution {
    public int findContentChildren(int[] g, int[] s) {
        Arrays.sort(g);
        Arrays.sort(s);
        int res = 0;
        int gCounter = 0;
        int sCounter = 0;
        while (gCounter < g.length && sCounter < s.length) {
            if (g[gCounter] > s[sCounter]) {
                sCounter++;
            }
            else {
                res++;
                gCounter++;
                sCounter++;
            }
        }
        return res;
    }
}