func findChampion(grid [][]int) int {
    for i, row := range grid {
        cnt := 0
        
        for _, cell := range row {
            cnt += cell
        }
        
        if cnt == len(grid) - 1 {
            return i
        }
    }
    
    return -1
}