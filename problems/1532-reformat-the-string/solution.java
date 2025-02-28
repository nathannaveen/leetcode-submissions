class Solution {
    public String reformat(String s) {
        List<Character> digits = new ArrayList<>();
        List<Character> letters = new ArrayList<>();
        String startsWith = "";
        String addToRes = "";
        StringBuilder res = new StringBuilder();
        int size = 0;

        for (char i : s.toCharArray()){
            if (Character.isDigit(i))
                digits.add(i);
            else
                letters.add(i);
        }
        System.out.println(letters.toString() + " " + digits.toString());

        if (digits.size() != letters.size() &&
                digits.size() - 1 != letters.size() &&
                digits.size() != letters.size() - 1)
            return "";

        if (digits.size() > letters.size()){
            addToRes += digits.get(digits.size() - 1);
            startsWith = "d";
            size = letters.size();
        }
        else if (digits.size() < letters.size()){
            addToRes += letters.get(letters.size() -1);
            startsWith = "l";
            size = digits.size();
        }
        else {
            size = letters.size();
            startsWith = "d";
        }

        for (int i = 0; i < size; i++) {
            if (startsWith.equals("d")){
                res.append(digits.get(i));
                res.append(letters.get(i));
            }
            else {
                res.append(letters.get(i));
                res.append(digits.get(i));
            }
        }
        
        res.append(addToRes);
        
        return res.toString();
    }
}