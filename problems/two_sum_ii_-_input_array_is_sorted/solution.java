public class Solution {
    public static int[] twoSum(int[] numbers, int target) {
//        System.out.println(Arrays.toString(numbers));
//        System.out.println(target);
        int left = 0, right = numbers.length-1;
        int[] result = new int[2];
        while (left < right){
            int sum = numbers[left] + numbers[right];
            if (sum == target){
                result[0] = left+1;
                result[1] = right+1;
                Arrays.sort(result);
                return result;
            }
            else if (sum < target){
                left ++;
            }
            else {
                right --;
            }
        }
        return null;
    }
}