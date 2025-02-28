class Solution {
    public List<String> commonChars(String[] A) {
        List<String> result = new ArrayList<>();
        int[] count = new int[26];
        Arrays.fill(count, Integer.MAX_VALUE);
        for (String str: A){
            int[] h = new int[26];
            for (char c : str.toCharArray()){
                ++h[c - 'a'];
            }

            for (int i = 0; i < 26; i++) {
                count[i] = Math.min(h[i], count[i]);
            }
        }
        for (char i = 'a'; i <= 'z'; i++) {
            while (count[i - 'a'] -- > 0){
                result.add("" + i);
            }
        }
        return result;
    }
}