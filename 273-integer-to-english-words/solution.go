func numberToWords(num int) string {
    if num == 0 {
        return "Zero"
    }
    
    beforeTwenty := []string{ "", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve", "Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen" }
    Tens := []string{ "Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
    others := []string{ "Thousand", "Million", "Billion" }
    
    arr := [][]int{}
    
    res := []string{}
    
    for num > 0 {
        n := []int{}
        
        for i := 0; i < 3 && num > 0; i++ {
            n = append([]int{ num % 10 }, n...)
            num /= 10
        }
        arr = append([][]int{ n }, arr...)
    }
    
    for i := 0; i < len(arr); i++ {
        done := false
        
        start := 0
        
        if len(arr[i]) == 3 && arr[i][0] != 0 {
            res = append(res, beforeTwenty[arr[i][0]])
            res = append(res, "Hundred")
            done = true
            start = 1
        }
        n := 0

        for j := start; j < len(arr[i]); j++ {
            n *= 10
            n += arr[i][j]
        }

        if n < 20 && n != 0 {
            res = append(res, beforeTwenty[n])
            done = true
        } else if (n >= 20) {
            onesDigit := n % 10
            n /= 10
            tensDigit := n

            res = append(res, Tens[tensDigit - 2])

            if onesDigit != 0 {
                res = append(res, beforeTwenty[onesDigit])
            }
            done = true
        }
        
        z := len(arr) - i - 2
        if z >= 0 && done {
            res = append(res, others[z])
        }
    }
    
    return strings.Join(res, " ")
}