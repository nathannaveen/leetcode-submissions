func findRotation(mat [][]int, target [][]int) bool {
    rotateOne, rotateTwo, rotateThree := true, true, true
    
    for i := 0; i < len(mat); i++ {
        for j := 0; j < len(mat); j++ {
            if mat[i][j] != target[len(target) - 1 - j][i] {
                rotateOne = false
            }
            if mat[i][j] != target[j][len(target) - 1 - i] {
                rotateTwo = false
            }
            if mat[i][j] != target[len(target) - 1 - i][len(target) - 1 - j] {
                rotateThree = false
            }
        }
    }
    
    return rotateOne || rotateTwo || rotateThree || reflect.DeepEqual(mat, target)
}