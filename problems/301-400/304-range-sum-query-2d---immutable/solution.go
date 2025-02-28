type NumMatrix struct {
    prefix [][]int
}


func Constructor(matrix [][]int) NumMatrix {
    prefix := [][]int{}

    for i := 0; i < len(matrix); i++ {
        prefix = append(prefix, []int{0})

        for j := 0; j < len(matrix[0]); j++ {
            prefix[i] = append(prefix[i], prefix[i][j] + matrix[i][j])
        }
    }

    return NumMatrix{prefix}
}


func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
    sum := 0

    for i := row1; i <= row2; i++ {
        sum += this.prefix[i][col2 + 1] - this.prefix[i][col1]
    }

    return sum
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */