func addStrings(num1 string, num2 string) string {
    carry := 0
    sum := ""

    for len(num1) > 0 || len(num2) > 0 {
        s := carry

        if len(num1) > 0 {
            n, _ := strconv.Atoi(string(num1[len(num1) - 1]))
            num1 = num1[:len(num1) - 1]
            s += n
        }
        if len(num2) > 0 {
            n, _ := strconv.Atoi(string(num2[len(num2) - 1]))
            num2 = num2[:len(num2) - 1]
            s += n
        }

        if s > 9 {
            carry = 1
            s -= 10
        } else {
            carry = 0
        }
        
        sum = strconv.Itoa(s) + sum
    }

    if carry != 0 {
        sum = strconv.Itoa(carry) + sum
    }

    return sum
}