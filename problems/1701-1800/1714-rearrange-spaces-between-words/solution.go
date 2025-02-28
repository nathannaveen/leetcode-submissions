func reorderSpaces(text string) string {
	spaceCounter := 0
	arr := []string{}
	str := ""
	res := ""
	for _, i := range text {
		if i == ' ' {
			if str != "" {
				arr = append(arr, str)
			}
			spaceCounter++
			str = ""
		} else {
			str += string(i)
		}
	}
	if str != "" {
		arr = append(arr, str)
	}
	space := ""

	if len(arr)-1 > 0 {
		for i := 0; i < spaceCounter/(len(arr)-1); i++ {
			space += " "
		}
		for i := 0; i < len(arr)-1; i++ {
			res += arr[i] + space
		}
		res += arr[len(arr)-1]
		for i := 0; i < spaceCounter%(len(arr)-1); i++ {
			res += " "
		}
	} else {
		for i := 0; i < spaceCounter; i++ {
			space += " "
		}
		res = strings.Trim(text, space) + space
	}
	return res
}