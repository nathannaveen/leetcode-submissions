impl Solution {
    pub fn get_concatenation(nums: Vec<i32>) -> Vec<i32> {
        let mut new_array = nums.clone();
        new_array.extend(&nums);
        new_array
    }
}

