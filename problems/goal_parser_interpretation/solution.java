class Solution {
    public String interpret(String command) {
        String result = "";
        int i = 0;
        while (i <= command.length() - 1){
            if (command.charAt(i) == '('){
                if (command.charAt(i + 1) == 'a'){
                    result += "al";
                    i += 4;
                }
                else {
                    result += "o";
                    i += 2;
                }
            }
            else {
                result += "G";
                i += 1;
            }
        }
        return result;
    }
}