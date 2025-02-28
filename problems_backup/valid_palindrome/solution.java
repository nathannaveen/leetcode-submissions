class Solution {
    public boolean isPalindrome(String s) {
        // ^[a-zA-Z]+$
        String[] theChars = s.split("");
        int left = 0;
        int right = theChars.length - 1;
        
        
        while(left<=right) {
            while (left < right && !theChars[left].matches("^[a-zA-Z0-9]*$")){
                left +=1;
            }
            
            while (left < right && !theChars[right].matches("^[a-zA-Z0-9]*$")){
                right -=1;
            }
            System.out.println(left + " " + right);
            
            if (!theChars[left].toLowerCase()
            .equals(theChars[right].toLowerCase())){
                return false;
            }
            
            
            right --;
            left ++;
        }
        return true;
        
    }
}