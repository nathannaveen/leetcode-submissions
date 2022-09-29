func minMovesToSeat(seats []int, students []int) int {
    sort.Ints(seats)
    sort.Ints(students)
    
    res := 0
    
    for i := 0; i < len(seats); i++ {
        res += int(math.Abs(float64(seats[i] - students[i])))
    }
    
    return res
}