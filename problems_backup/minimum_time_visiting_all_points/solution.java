class Solution {
    public int minTimeToVisitAllPoints(int[][] points) {
        long sum = 0;
        for (int i = 0; i < points.length - 1; i++) {
            long x = points[i][0];
            long y = points[i][1];
            long tx = Math.abs(points[i+1][0] - x);
            long ty = Math.abs(points[i+1][1] - y);
            while (tx != ty){
                if (tx > ty){
                    tx -= 1;
                }
                if (ty > tx){
                    ty -= 1;
                }
            }
            if (points[i+1][0] - x > 0){
                if (points[i+1][1] - y > 0){
                    x += tx;
                    y += ty;
                }
                else {
                    x += tx;
                    y -= ty;
                }
            }
            else {
                if (points[i+1][1] - y > 0){
                    x -= tx;
                    y += ty;
                }
                else {
                    x -= tx;
                    y -= ty;
                }

            }
            sum += tx;
            if (x == points[i+1][0]){
                sum += Math.abs(points[i+1][1] - y);
            }
            else {

                sum += Math.abs(points[i+1][0] - x);
            }
        }
        return (int) sum;
    }
}