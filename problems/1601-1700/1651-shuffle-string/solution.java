class Solution {
    public String restoreString(String s, int[] indices) {
        String[] h = s.split("");
        String[] u = new String[s.length()];
        for (int i = 0; i < h.length; i++) {
            u[indices[i]] = h[i];
        }
        s = "";
        for (int i = 0; i < u.length; i++) {
            s += u[i];
        }
        return s;
    }
}