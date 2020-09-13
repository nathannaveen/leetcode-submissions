import java.util.Arrays;
public class Solution {
    public static int[] twoSum(int[] numbers, int target) {
        for (int i = 0; i < numbers.length; i++) {
            for (int j = i+1; j < numbers.length; j++) {
                if (numbers[i] + numbers[j] == target){
                    int[] result = new int[]{i,j};
                    Arrays.sort(result);
                    return result;

                }
            }
        }
        return null;
    }
}