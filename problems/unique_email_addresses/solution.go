func numUniqueEmails(emails []string) int {
	m := make(map[string]int)
	res := 0

	for _, email := range emails {
		end := strings.Index(email, "@")
		email = strings.ReplaceAll(email[:end], ".", "") + email[end:]
		start := strings.Index(email, "+")
		end = strings.Index(email, "@")
		if start != -1 { email = email[:start] + email[end:] }
		if m[email] == 0 { res, m[email] = res + 1, 1 }
	}

	return res
}
