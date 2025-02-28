class Solution {
    public String[] uncommonFromSentences(String A, String B) {
        HashMap<String, Integer> h = new HashMap<>();
        List<String> n = new ArrayList<>();

        for (String i : A.split(" ")){
            if (h.containsKey(i)){
                int g = h.get(i);
                h.replace(i, g + 1);
            }
            else {
                h.put(i, 1);
            }
        }

        for (String i : B.split(" ")){
            if (h.containsKey(i)){
                int g = h.get(i);
                h.replace(i, g + 1);
            }
            else {
                h.put(i, 1);
            }
        }

        for (String i : h.keySet()){
            if (h.get(i) == 1) n.add(i);
        }

        String[] g = new String[n.size()];

        for (int i = 0; i < g.length; i++) {
            g[i] = n.get(i);
        }
        
        return g;
    }
}