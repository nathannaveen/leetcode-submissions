func isBoomerang(points [][]int) bool {
    a := (points[0][0] - points[2][0]) * (points[0][1] - points[1][1])
    b := (points[0][0] - points[1][0]) * (points[0][1] - points[2][1])
    return a != b
}