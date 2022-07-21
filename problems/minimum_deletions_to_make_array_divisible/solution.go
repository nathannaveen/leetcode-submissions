func minOperations(nums []int, numsDivide []int) int {
    n := arrayGCD(numsDivide, 0)
    
    sort.Ints(nums)
    
    for i := 0; i < len(nums); i++ {
        if n % nums[i] == 0 {
            return i
        }
    }
    
    return -1
}

func arrayGCD(arr []int, i int) int {
    if i == len(arr) - 1 {
        return arr[i]
    }
    
    a := arr[i]
    b := arrayGCD(arr, i + 1)
    return gcd(a, b)
}

func gcd(a, b int) int {
    if a == 0 {
        return b
    }
    return gcd(b % a, a)
}