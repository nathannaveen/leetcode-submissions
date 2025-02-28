func romanToInt(s string) int {
	res := 0

	for i := len(s) - 1; i >= 0; i-- {
		shouldSubtract := false

		if i != len(s)-1 {
			if (string(s[i+1]) == "V" || string(s[i+1]) == "X") && string(s[i]) == "I" {
				shouldSubtract = true
				res -= 1
			} else if (string(s[i+1]) == "L" || string(s[i+1]) == "C") && string(s[i]) == "X" {
				shouldSubtract = true
				res -= 10
			} else if (string(s[i+1]) == "M" || string(s[i+1]) == "D") && string(s[i]) == "C" {
				shouldSubtract = true
				res -= 100
			}
		}

		if !shouldSubtract {
			switch string(s[i]) {
			case "I":
				res += 1
			case "X":
				res += 10
			case "C":
				res += 100
			case "V":
				res += 5
			case "L":
				res += 50
			case "D":
				res += 500
			case "M":
				res += 1000
			}

		}
	}
	return res
}