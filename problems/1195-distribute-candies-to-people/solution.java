class Solution {
    public int[] distributeCandies(int candies, int num_people) {
        int[] h = new int[num_people];
        int people  = 0;
        int counter = 1;
        while (candies > 0){
            if (candies >= counter){
                h[people] += counter;
                candies -= counter;
            }
            else {
                h[people] += candies;
                candies = 0;
            }
            if (people == num_people - 1){
                people = -1;
            }
            people++;
            counter ++;
        }
        return h;
    }
}