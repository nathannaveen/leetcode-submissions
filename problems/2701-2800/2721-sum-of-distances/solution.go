// func distance(nums []int) []int64 {
//     indexes := map[int] []int{}
//     arr := make([]int64, len(nums))
    
//     for i := 0; i < len(nums); i++ {
//         z := indexes[nums[i]]
//         for j := 0; j < len(z); j++ {
//             t := int64(i - z[j])
//             arr[i] += t
//             arr[z[j]] += t
//         }
        
//         indexes[nums[i]] = append(indexes[nums[i]], i)
//     }
    
//     return arr
// }

func distance(nums []int) []int64 {
    indices := make(map[int][]int)
    for i, num := range nums {
        indices[num] = append(indices[num], i)
    }
    res := make([]int64, len(nums))
    for k := range indices {
        totalSum := 0
        for _, i := range indices[k] {
            totalSum += i
        }
        preSum := 0
        for i, idx := range indices[k] {
            dist := idx*i - preSum + (totalSum-preSum-idx-idx*(len(indices[k])-i-1))
            res[idx] = int64(dist)
            preSum += idx
        }
    }
    return res
}