class Solution {
    public List<String> buildArray(int[] target, int n) {
        List<String> res = new ArrayList<>();
        int targetCounter = 0;
        for (int i = 1; i < n+1; i++) {
            if (target[targetCounter] != i){
                res.add("Push");
                res.add("Pop");
            }
            else if (target[targetCounter] == i){
                res.add("Push");
                targetCounter ++;
            }
            if (targetCounter == target.length){
                break;
            }
        }
        return res;
    }
}