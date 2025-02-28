class Solution {
    public String addBinary(String a, String b) {
        StringBuilder string = new StringBuilder();
        int alength = a.length()-1;
        int blength = b.length()-1;
        int carry = 0;
        while(alength>=0 || blength>=0){
            int s = carry;
            if(alength>=0){
               s += a.charAt(alength--)-48;
            }
            if(blength>=0){
               s += b.charAt(blength--)-48;
            }
            string.insert(0 ,s % 2);
            carry = s / 2;  
        }
        if(carry != 0){
            string.insert(0,carry);
        }
        return string.toString();  
    }
}