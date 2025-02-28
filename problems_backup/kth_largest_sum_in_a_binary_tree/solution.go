/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthLargestLevelSum(root *TreeNode, k int) int64 {
    sums := []int64{}
    queue := []*TreeNode{ root }

    for len(queue) > 0 {
        n := len(queue)
        sum := int64(0)
        
        for i := 0; i < n; i++ {
            pop := queue[0]
            queue = queue[1:]

            if pop == nil {
                continue
            }

            sum += int64(pop.Val)
            queue = append(queue, pop.Right, pop.Left)
        }
        sums = append(sums, sum)
    }

    sort.Slice(sums, func(i, j int) bool {
        return sums[i] > sums[j]
    })

    if len(sums) < k + 1 {
        return -1
    }

    return sums[k - 1]
}