class Solution {
    public String reformatNumber(String number) {
        StringBuilder h = new StringBuilder();
        StringBuilder phoneNumber = new StringBuilder();
//        Set<Character> num = new HashSet<>();
//
//        num.add('0'); num.add('1'); num.add('2'); num.add('3'); num.add('4');
//        num.add('5'); num.add('6'); num.add('7'); num.add('8'); num.add('9');

        for (int i = 0; i < number.length(); i++) {
            if ((int) number.charAt(i) <= 57 && (int) number.charAt(i) >= 48) {
                h.append(number.charAt(i));
            }
        }

        if (h.length() == 2){
            return h.toString();
        }

        int i = h.length() / 3;
        if (h.length() % 3 == 1){
            i -= 1;
        }
        int counter = 0;

        for (int j = 0; j < i; j++) {
            phoneNumber.append(h.charAt(counter));
            phoneNumber.append(h.charAt(counter + 1));
            phoneNumber.append(h.charAt(counter + 2));
            if (j != i - 1 && h.length() % 3 == 0) {
                phoneNumber.append("-");
            }
            else if (h.length() % 3 != 0){
                phoneNumber.append("-");
            }
            counter += 3;
        }

        if (h.length() % 3 == 1){
            phoneNumber.append(h.charAt(counter));
            phoneNumber.append(h.charAt(counter + 1));
            phoneNumber.append("-");
            phoneNumber.append(h.charAt(counter + 2));
            phoneNumber.append(h.charAt(counter + 3));
        }

        if (h.length() % 3 == 2){
            phoneNumber.append(h.charAt(counter));
            phoneNumber.append(h.charAt(counter + 1));
        }
        return phoneNumber.toString();
    }
}