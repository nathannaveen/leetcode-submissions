class Solution {
    public List<String> subdomainVisits(String[] cpdomains) {
        HashMap<String, Integer> h = new HashMap<>();
        for (int i = 0; i < cpdomains.length; i++) {
            String[] v = cpdomains[i].split(" ");
            String word = v[1];
            int number = Integer.parseInt(v[0]);
            int start = 0;
            for (int j = 0; j < word.length(); j++) {
                if (word.charAt(j) == '.') {
                    String theKey = word.substring(start, word.length());
                    if (h.containsKey(theKey)){
                        int newNumber = h.get(theKey) + number;
                        h.replace(theKey, newNumber);
                    }
                    else {
                        h.put(theKey, number);
                    }
                    start = j + 1;
                }
            }
            if (h.containsKey(word.substring(start, word.length()))) {
                int newNumber = h.get(word.substring(start, word.length())) + number;
                h.replace(word.substring(start, word.length()), newNumber);
            }
            else {
                h.put(word.substring(start, word.length()), number);
            }
        }
        List<String> result = new ArrayList<>();

        for (String i : h.keySet()){
            result.add(Integer.toString(h.get(i)) + " " + i);
        }

        return result;
    }
}