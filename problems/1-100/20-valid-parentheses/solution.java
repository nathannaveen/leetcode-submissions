class Solution {
    public boolean isValid(String s) {
        Stack<Character> h = new Stack<>();
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) == '{' || s.charAt(i) == '(' || s.charAt(i) == '['){
                h.add(s.charAt(i));
            }
            else {
                if (s.charAt(i) == '}'){
                    if (h.size() == 0 || h.size() > 0 && h.peek() != '{'){
                        return false;
                    }
                    else if (h.size()>0)
                        h.pop();
                }
                if (s.charAt(i) == ']'){
                    if (h.size() == 0 || h.size() > 0 && h.peek() != '['){
                        return false;
                    }
                    else if (h.size()>0)
                        h.pop();
                }
                if (s.charAt(i) == ')'){
                    if (h.size() == 0 || h.size() > 0 && h.peek() != '('){
                        return false;
                    }
                    else if (h.size()>0)
                        h.pop();
                }
            }
        }
        return h.size() == 0;
    }
}