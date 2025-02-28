class Solution {
    public boolean isNStraightHand(int[] hand, int W) {

        if (hand.length == W){
            return IfHandLengthEqualsW(hand);
        }
        if (hand.length % W == 0){
            Arrays.sort(hand);
            boolean[] h = new boolean[hand.length];
            Arrays.fill(h, false);
            int[][] g = new int[h.length / W][W];
            for (int i = 0; i < hand.length / W; i++) {
                int counter = 0;
                for (int j = 0; j < hand.length; j++) {
                    if (counter == 0){
                        if (!h[j]){
                            g[i][counter] = hand[j];
                            h[j] = true;
                            counter ++;
                        }
                    }
                    else {

                        if (counter >= W){
                            break;
                        }
                        if (!h[j]){
                            if (hand[j] == g[i][counter - 1] + 1){
                                g[i][counter] = hand[j];
                                h[j] = true;
                                counter ++;
                            }
                        }
                    }
                }
            }
            if (PrintingAndChecking(g)) return false;
            return true;
        }
        else{
            return false;
        }
    }

    private boolean PrintingAndChecking(int[][] g) {
        for (int k = 0; k < g.length; k++) {
            for (int l = 0; l < g[k].length; l++) {
                if (g[k][l] == 0){
                    return true;
                }
                System.out.println(g[k][l]);
            }
            System.out.println("\n");
        }
        return false;
    }

    private boolean IfHandLengthEqualsW(int[] hand) {
        Arrays.sort(hand);
        System.out.println(Arrays.toString(hand));
        int oldNumber = hand[0];
        boolean g = true;
        for (int i = 1; i < hand.length; i++) {
            if (oldNumber != hand[i] - 1){
                g = false;
                break;
            }
            oldNumber = hand[i];

        }
        return g;
    }
}