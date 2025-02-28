/*
// Definition for Employee.
class Employee {
    public int id;
    public int importance;
    public List<Integer> subordinates;
};
*/

class Solution {
    public int getImportance(List<Employee> employees, int id) {
        int sum = 0;
        Stack<Integer> h = new Stack<>();

        for (Employee e : employees){
            if (e.id == id){
                for (int i = 0; i < e.subordinates.size(); i++) {
                    h.push(e.subordinates.get(i));
                }
                sum += e.importance;
            }
        }
        while (!h.isEmpty()){
            int g = h.pop();
            for (Employee e : employees){
                if (e.id == g){
                    for (int i = 0; i < e.subordinates.size(); i++) {
                        h.push(e.subordinates.get(i));
                    }
                    sum += e.importance;
                }
            }
        }
        return sum;
    }
}