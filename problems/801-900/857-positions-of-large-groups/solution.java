class Solution {
    public List<List<Integer>> largeGroupPositions(String s) {
        List<List<Integer>> list = new ArrayList<>();
        int start = 0;
        char letter = 0;
        int counter = 0;
        for (int i = 0; i < s.length(); i++) {
            if (s.charAt(i) != letter){
                GreaterThanOrEqualToThree(list, start, counter, i);
                letter = s.charAt(i);
                counter = 1;
                start= i;
            }
            else {
                counter ++;
            }
        }
        GreaterThanOrEqualToThree(list, start, counter, s.length());
        return list;
    }

    private void GreaterThanOrEqualToThree(List<List<Integer>> list, int start, int counter, int i) {
        if (counter >= 3) {
            List<Integer> h = new ArrayList<>();
            h.add(start);
            h.add(i - 1);
            list.add(h);
        }
    }
}