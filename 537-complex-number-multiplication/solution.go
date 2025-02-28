func complexNumberMultiply(num1 string, num2 string) string {
    real1, imag1 := split(num1)
    real2, imag2 := split(num2)
    a := real1 * real2
    b := imag1 * imag2
    c := real2 * imag1 + real1 * imag2

    return strconv.Itoa(a - b) + "+" + strconv.Itoa(c) + "i"
}

func split(s string) (int, int) {
    s = s[:len(s) - 1] // remove i because num1 and num2 always have it
    x := strings.Split(s, "+")
    real, _ := strconv.Atoi(x[0])
    imaginary, _ := strconv.Atoi(x[1])

    return real, imaginary
}