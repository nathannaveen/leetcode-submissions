class Solution {
    public List<String> findAndReplacePattern(String[] words, String pattern) {
        List<String> h = new ArrayList<>();
        HashMap<Character, Integer> g = new HashMap<>();
        int patternResult = 0;
        char m = '0';
        int counter = 0;
        for (int i = 0; i < pattern.length(); i++) {
            if (pattern.charAt(i) != m){
                m = pattern.charAt(i);
                patternResult += counter;
                if (g.containsKey(m)){
                    counter = g.get(m);
                }
                else {
                    counter++;
                    g.put(m, counter);
                }
            }
        }
        patternResult += counter;
        for (int i = 0; i < words.length; i++) {
            HashMap<Character, Integer> c = new HashMap<>();
            int result = 0;
            m = '0';
            counter = 0;
            for (int j = 0; j < words[i].length(); j++) {
                if (words[i].charAt(j) != m){
                    result += counter;
                    m = words[i].charAt(j);
                    if (c.containsKey(m)){
                        counter = c.get(m);
                    }
                    else{
                        counter++;
                        c.put(m , counter);
                    }
                }
            }
            result += counter;
            if (result == patternResult){
                h.add(words[i]);
            }
        }
        return h;
    }
}