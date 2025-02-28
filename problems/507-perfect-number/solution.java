class Solution {
    public boolean checkPerfectNumber(int num) {
        int sum = 0;
        if (num == 0){
            return false;
        }
        System.out.println(num);
        for (int i = 1; i <= num/2; i++) {
            if (num % i == 0){
                sum += i;
            }
        }
        if (sum == num) return true;
        return false;
    }
}