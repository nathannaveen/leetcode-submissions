class Solution {
    public boolean isHappy(int n) {
        HashMap<Integer, Integer> h = new HashMap<>();
        return internalisHappy(n, h);

    }

    public boolean internalisHappy(int n, HashMap<Integer, Integer> h) {
        if (h.containsKey(n)) {
            return false;
        }
        h.put(n, 1);
        int g = n;
        int sum = 0;
        while (g > 0) {
            sum += (g % 10) * (g % 10);
            g /= 10;
        }
        if (sum == 1) {
            return true;
        }
        return internalisHappy(sum, h);
    }
}