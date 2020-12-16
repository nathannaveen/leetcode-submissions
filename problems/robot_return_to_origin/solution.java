class Solution {
    public boolean judgeCircle(String moves) {
        int L = 0;
        int R = 0;
        int U = 0;
        int D = 0;
        for (char i : moves.toCharArray()){
            switch (i){
                case 'L':
                    L += 1;
                    break;
                case 'R':
                    R += 1;
                    break;
                case 'U':
                    U += 1;
                    break;
                case 'D':
                    D += 1;
                    break;
            }
        }
        if (L == R && U == D) return true;
        else return false;
    }
}