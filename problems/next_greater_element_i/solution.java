class Solution {
    public int[] nextGreaterElement(int[] nums1, int[] nums2) {
        int[] h = new int[nums1.length];
        
        for (int i = 0; i < nums1.length; i++){
            boolean filled = false;
            for (int j = 0; j < nums2.length; j++) {
                if (nums1[i] == nums2[j]){
                    for (int k = j + 1; k < nums2.length; k++) {
                        if (nums1[i] < nums2[k]){
                            h[i] = nums2[k];
                            filled = true;
                            break;
                        }
                    }
                    
                    if (filled){
                        break;
                    }
                    h[i] = -1;
                    break;
                }
            }
        }
        return h;
    }
}