func reformatDate(date string) string {
	month := map[string]string{
		"Jan": "01", "Feb": "02", "Mar": "03", "Apr": "04", "May": "05", "Jun": "06", 
		"Jul": "07", "Aug": "08", "Sep": "09", "Oct": "10", "Nov": "11", "Dec": "12",
	}
	if date[1] == 't' || date[1] == 'r' || date[1] == 's' || date[1] == 'n' {
		date = "0" + date
	}
	res := date[9:] + "-" + month[date[5:8]] + "-" + date[:2]
	return res
}
