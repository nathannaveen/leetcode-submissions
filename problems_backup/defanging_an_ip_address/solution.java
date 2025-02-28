class Solution {
    public String defangIPaddr(String address) {
        StringBuilder h = new StringBuilder();
        for (char i : address.toCharArray()){
            if (i =='.'){
                h.append("[.]");
            }
            else {
                h.append(i);
            }
        }
        return h.toString();
    }
}