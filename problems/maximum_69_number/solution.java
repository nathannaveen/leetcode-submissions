class Solution {
    public int maximum69Number (int num) {
        String[] h = Integer.toString(num).split("");
        for (int i = 0; i < h.length; i++) {
            if (h[i].equals("6")){
                h[i] = "9";
                break;
            }
        }
        num = 0;
        for (int i = 0; i < h.length; i++) {
            num *= 10;
            num += Integer.parseInt(h[i]);
        }
        return num;
    }
}