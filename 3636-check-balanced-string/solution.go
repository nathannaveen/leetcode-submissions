func isBalanced(num string) bool {
    a, b := 0, 0
    for i := 0; i < len(num); i++ {
        x := int(num[i] - '0')
        if i % 2 == 0 {
            a += x
        } else {
            b += x
        }
    }

    return a == b
}
