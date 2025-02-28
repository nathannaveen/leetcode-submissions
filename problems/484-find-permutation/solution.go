func findPermutation(s string) []int {
	res := make([]int, len(s)+1)

	for i := 1; i <= len(s) + 1; i++ { res[i - 1] = i }

	for i := 0; i < len(s); i++ {
		if s[i] == 'D' {
			temp := i
			for temp < len(s) && s[temp] == 'D' { temp++ }
			res = append(res[: i], append(reverseArray(res[i : temp + 1]), res[temp + 1 :]...)...)
			i = temp
		}
	}
	return res
}

func reverseArray(res []int) []int {
	left, right := 0, len(res) - 1
	for left < right {
		res[left], res[right] = res[right], res[left]
		left, right = left + 1, right - 1
	}
	return res
}