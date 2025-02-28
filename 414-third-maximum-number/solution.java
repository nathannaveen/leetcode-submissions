class Solution {
    public int thirdMax(int[] nums) {
        Arrays.sort(nums);
        int c = 0;
        int z = nums.length -1;
        while(z > c){
            int t = nums[c];
            nums[c] = nums[z];
            nums[z] = t;
            z--;
            c++;
        }
        
        int max = nums[0];
        int counter = 0;
        for(int i = 1;i < nums.length; i++){
            if(nums[i] != nums[i-1]){
                counter++;
            }

             if(counter == 2){
                return nums[i];
            }
        }
        return max;
    }
}