func getHint(secret string, guess string) string {
    m := map[byte] int {}
    bulls, cows := 0, 0
    

    for i := 0; i < len(secret); i++ {
        m[secret[i]]++
        if guess[i] == secret[i] {
            bulls++
            m[secret[i]]--
        }
    }

    for i := 0; i < len(guess); i++ {
        if guess[i] != secret[i] && m[guess[i]] > 0 {
            cows++
            m[guess[i]]--
        }
    }

    return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}