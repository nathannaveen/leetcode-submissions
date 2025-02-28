class Solution {
    public int countCharacters(String[] words, String chars) {
        char[] h = chars.toCharArray();
        Arrays.sort(h);
        int counter = 0;

        for (int i = 0; i < words.length; i++) {
            char[] n = words[i].toCharArray();
            Arrays.sort(n);
            int nCounter = 0;
            int hCounter = 0;
            for (int j = 0; j < h.length; j++) {
                if (h[j] == n[nCounter]){
                    System.out.println("taco");
                    nCounter ++;
                }
                if (nCounter == n.length){
                    break;
                }
            }
            if (nCounter == n.length){
                counter += nCounter;
            }
        }
        return counter;
    }
}