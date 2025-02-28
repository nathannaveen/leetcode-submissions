class Solution {
    public void merge(int[] nums1, int m, int[] nums2, int n) {
        int counter = 0;
        for (int i = m; i < nums1.length; i++) {
            nums1[i] = nums2[counter];
            counter ++;
        }
        for (int i = nums1.length - 1; i >= 1; i--) {
            for (int j = i; j >= 0; j--) {
                if (nums1[j] > nums1[i]){
                    int g = nums1[i];
                    nums1[i] = nums1[j];
                    nums1[j] = g;
                }
            }
        }
    }
}