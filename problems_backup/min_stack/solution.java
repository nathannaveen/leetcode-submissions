import java.util.ArrayList;
import java.util.List;

class MinStack {

    /** initialize your data structure here. */
    List<Integer> h;
    public MinStack() {
        h = new ArrayList<>();
    }

    public void push(int x) {
        h.add(x);
    }

    public void pop() {
        h.remove(h.size() - 1);
    }

    public int top() {
        return h.get(h.size() - 1);
    }

    public int getMin() {
        int min = Integer.MAX_VALUE;
        for (int g : h){
            if (g < min){
                min = g;
            }
        }
        return min;
    }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack obj = new MinStack();
 * obj.push(x);
 * obj.pop();
 * int param_3 = obj.top();
 * int param_4 = obj.getMin();
 */