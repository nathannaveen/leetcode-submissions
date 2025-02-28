func countStudents(students []int, sandwiches []int) int {
	kidGotSandwich := true

	for kidGotSandwich {
		kidGotSandwich = false
		for i := 0; i < len(students); i++ {
			pop := students[0]
			students = students[1:]
			if pop == sandwiches[0] {
				kidGotSandwich = true
				sandwiches = sandwiches[1:]
				break
			} else {
				students = append(students, pop)
			}
		}
	}
	return len(students)
}