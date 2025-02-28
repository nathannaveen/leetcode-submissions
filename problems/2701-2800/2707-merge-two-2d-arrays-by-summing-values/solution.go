func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
    var i, j int 
    res := [][]int{}

    for i < len(nums1) || j < len(nums2) {
        oneId := 1001
        twoId := 1001

        if i != len(nums1) {
            oneId = nums1[i][0]
        }
        if j != len(nums2) {
            twoId = nums2[j][0]
        }

        if oneId < twoId {
            res = append(res, nums1[i])
            i++
        } else if twoId < oneId {
            res = append(res, nums2[j])
            j++
        } else {
            res = append(res, []int{ nums1[i][0], nums1[i][1] + nums2[j][1] })
            i++
            j++
        }
    }

    return res
}