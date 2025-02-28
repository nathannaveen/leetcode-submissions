class Solution {
    public List<List<Integer>> groupThePeople(int[] groupSizes) {
        HashMap<Integer, Integer> map = new HashMap<>();

        List<List<Integer>> result = new ArrayList<>();

        for (int i = 0; i < groupSizes.length; i++) {
            map.put(i, groupSizes[i]);
        }

        for (int groupSize : groupSizes) {
            List<Integer> innerList = new ArrayList<>();
            boolean contains = false;
            for (int key : map.keySet()) {
                if (map.get(key) == groupSize && map.get(key) != -1) {
                    contains = true;
                    innerList.add(key);
                    map.replace(key, -1);
                }
                if (innerList.size() == groupSize) break;
            }
            if (contains) result.add(innerList);
        }
        return result;
    }
}