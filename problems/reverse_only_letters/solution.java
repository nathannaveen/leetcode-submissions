class Solution {
    public String reverseOnlyLetters(String S) {
        int left = 0;
        int right = S.length() - 1;
        char[] h = S.toCharArray();
        Set<Character> letters = new HashSet<>();

        letters.add('q'); letters.add('w'); letters.add('e'); letters.add('r'); letters.add('t'); // lowercase
        letters.add('y'); letters.add('u'); letters.add('i'); letters.add('o'); letters.add('p');
        letters.add('a'); letters.add('s'); letters.add('d'); letters.add('f'); letters.add('g');
        letters.add('h'); letters.add('j'); letters.add('k'); letters.add('l'); letters.add('z'); 
        letters.add('x'); letters.add('c'); letters.add('v'); letters.add('b'); letters.add('n'); 
        letters.add('m');
        
        letters.add('Q'); letters.add('W'); letters.add('E'); letters.add('R'); letters.add('T'); // Uppercase
        letters.add('Y'); letters.add('U'); letters.add('I'); letters.add('O'); letters.add('P');
        letters.add('A'); letters.add('S'); letters.add('D'); letters.add('F'); letters.add('G');
        letters.add('H'); letters.add('J'); letters.add('K'); letters.add('L'); letters.add('Z'); 
        letters.add('X'); letters.add('C'); letters.add('V'); letters.add('B'); letters.add('N'); 
        letters.add('M');

        while (left < right){
            if (!letters.contains(h[left])){
                left ++;
            }
            if (!letters.contains(h[right])){
                right --;
            }
            if (letters.contains(h[right]) && letters.contains(h[left])) {
                char y = h[right];
                h[right] = h[left];
                h[left] = y;
                left++;
                right--;
            }
        }
        String n = "";
        for (char i : h){
            n += i;
        }
        return n;
    }
}