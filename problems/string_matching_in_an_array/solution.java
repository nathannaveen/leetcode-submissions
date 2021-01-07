class Solution {
    public List<String> stringMatching(String[] words) {
        List<String> h = new ArrayList<>();
        for (int i = 0; i < words.length; i++) {
            for (int j = i + 1; j < words.length; j++) {
                if (words[i].contains(words[j]) && !h.contains(words[j])){
                    h.add(words[j]);
                }
                if (words[j].contains(words[i]) && !h.contains(words[i])){
                    h.add(words[i]);
                }
            }
        }
        return h;
    }
}