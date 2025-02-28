func numUniqueEmails(emails []string) int {
	m := make(map[string]int)

	for _, email := range emails {
		end := strings.Index(email, "@")
		email = strings.ReplaceAll(email[:end], ".", "") + email[end:]
		start := strings.Index(email, "+")
		end = strings.Index(email, "@")
		if start != -1 { 
            email = email[:start] + email[end:] 
        }
		m[email]++
	}

    return len(m)
}
