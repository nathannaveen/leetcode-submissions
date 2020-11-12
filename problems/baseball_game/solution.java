class Solution {
    public int calPoints(String[] ops) {
        int sum = 0;
        Stack<Integer> h = new Stack<>();
        for (int i = 0; i < ops.length; i++) {
            switch (ops[i]){
                case "+":
                    h.push(h.peek() + h.get(h.size()-2));
                    break;
                case "D":
                    h.push(h.peek() * 2);
                    break;
                case "C":
                    h.pop();
                    break;
                default:
                    h.push(Integer.parseInt(ops[i]));
                    break;
            }
        }
        for (int i = h.size()-1; i > -1; i--) {
            sum += h.peek();
            h.pop();
        }
        return sum;
    }
}