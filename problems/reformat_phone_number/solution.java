class Solution {
    public String reformatNumber(String number) {
        StringBuilder h = new StringBuilder();
        StringBuilder phoneNumber = new StringBuilder();

        for (int i = 0; i < number.length(); i++) {
            if (Character.isDigit(number.charAt(i))){
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