class Solution {
    public boolean canConstruct(String ransomNote, String magazine) {
        int[] chars = new int[26];

        for (int i = 0; i < Math.max(ransomNote.length(), magazine.length()); i++) {
            if (magazine.length() > i){
                chars[(int) (magazine.charAt(i) - 'a')] --;
                System.out.println((int) (magazine.charAt(i) - 'a'));
            }
            if (ransomNote.length() > i){
                chars[(int) (ransomNote.charAt(i) - 'a')] ++;
                System.out.println((int) (ransomNote.charAt(i) - 'a'));
            }
        }
        for (int i : chars){
            if (i > 0) return false;
        }
        return true;
//         if (ransomNote.length() == 0){
//             return true;
//         }

//         String[] note = ransomNote.split("");
//         String[] mag = magazine.split("");
//         Arrays.sort(note);
//         Arrays.sort(mag);

//         int magCounter = 0;
//         for (String s : note) {
//             if (magCounter == mag.length){
//                 return false;
//             }
//             while (!mag[magCounter].equals(s)) {
//                 if (magCounter == mag.length - 1) {
//                     return false;
//                 }
//                 if (!mag[magCounter].equals(s)) {
//                     magCounter++;
//                 }
//             }
//             magCounter ++;
//         }
//         return true;
    }
}