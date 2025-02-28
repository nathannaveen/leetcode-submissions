class Solution {
    public int maxPower(String s) {
        int power = 0;
        char letter = 0;
        int smallPower = 0;
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == letter){
                smallPower ++;
            }
            else {
                letter = s.charAt(i);
                smallPower = 1;
            }
            
            if (smallPower > power){
                power = smallPower;
            }
        }
        return power;
    }
}