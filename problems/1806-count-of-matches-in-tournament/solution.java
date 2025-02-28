class Solution {
    public int numberOfMatches(int n) {
        int sum = 0;
        int newN = 0;
        while (n > 1){
            if (n % 2 == 1) newN = (n / 2) + 1;
            else newN = n / 2;
            sum += n / 2;
            n = newN;
        }
        return sum;
    }
}