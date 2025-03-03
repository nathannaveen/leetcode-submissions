func minSwaps(nums []int) int {
    numOnes := 0
    numOnesWindow := 0
    
    for i := 0; i < len(nums); i++ {
        numOnes += nums[i]
    }
    
    for i := 0; i < numOnes; i++ {
        numOnesWindow += nums[i]
    }
    
    maxInWindow := numOnesWindow
    
    for i := numOnes; i < len(nums) * 2; i++ {
        numOnesWindow -= nums[(i - numOnes) % len(nums)]
        numOnesWindow += nums[i % len(nums)]
        maxInWindow = max(maxInWindow, numOnesWindow)
    }
    
    return numOnes - maxInWindow
}

func max(a, b int) int {
    if a > b { return a }
    return b
}