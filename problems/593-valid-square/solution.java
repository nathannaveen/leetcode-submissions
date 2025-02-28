class Solution {
    public boolean validSquare(int[] p1, int[] p2, int[] p3, int[] p4) {
        int count =0;
        Double d1 = Math.pow((p1[0] - p2[0]),2)+ Math.pow((p1[1] - p2[1]),2);
        Double d2 = Math.pow((p2[0] - p3[0]),2)+ Math.pow((p2[1] - p3[1]),2);
        Double d3 = Math.pow((p3[0] - p4[0]),2)+ Math.pow((p3[1] - p4[1]),2);
        Double d4 = Math.pow((p4[0] - p1[0]),2)+ Math.pow((p4[1] - p1[1]),2);
        Double d5 = Math.pow((p1[0] - p3[0]),2)+ Math.pow((p1[1] - p3[1]),2);
        Double d6 = Math.pow((p2[0] - p4[0]),2)+ Math.pow((p2[1] - p4[1]),2);
        List<Double> arr = Arrays.asList(d1,d2,d3,d4,d5,d6);
        Collections.sort(arr);
        System.out.println(arr.toString());
        for(int i=1; i <6; i++){
            if(arr.get(i).equals(arr.get(i-1)) && !arr.get(i).equals(0.0))
                count++;
        }
        
        if(count ==4)
            return true;
        
        return false;
    }
    //     float distP1P2 = (float) Math.sqrt(Math.pow((p1[0] - p2[0]), 2) + Math.pow(p1[1] - p2[1], 2));
    //     float distP1P3 = (float) Math.sqrt(Math.pow((p1[0] - p3[0]), 2) + Math.pow(p1[1] - p3[1], 2));
    //     float distP1P4 = (float) Math.sqrt(Math.pow((p1[0] - p4[0]), 2) + Math.pow(p1[1] - p4[1], 2));
        
    //     if (distP1P2 > distP1P3 && distP1P2 > distP1P4){
    //         float distP4P3 = (float) Math.sqrt(Math.pow((p4[0] - p3[0]), 2) + Math.pow(p4[1] - p3[1], 2));
    //         if (distP1P2 == distP4P3) return true;
    //         else return false;
    //     }
    //     if (distP1P3 > distP1P2 && distP1P3 > distP1P4){
    //         float distP2P4 = (float) Math.sqrt(Math.pow((p2[0] - p4[0]), 2) + Math.pow(p2[1] - p4[1], 2));
    //         if (distP2P4 == distP1P3) return true;
    //         else return false;
    //     }
    //     if (distP1P4 > distP1P2 && distP1P4 > distP1P3){
    //         float distP2P3 = (float) Math.sqrt(Math.pow((p2[0] - p3[0]), 2) + Math.pow(p2[1] - p3[1], 2));
    //         if (distP1P4 == distP2P3) return true;
    //         else return false;
    //     }
        
    //     return false;
    // }
}