func canPlaceFlowers(flowerbed []int, n int) bool {
    flowerbed = append([]int{0}, append(flowerbed, 0)...)
    
    counter := 0
    for i := 1; i < len(flowerbed) - 1; i++ {
        if flowerbed[i - 1] + flowerbed[i] + flowerbed[i + 1] == 0 {
            i++
            counter++
        }
    }
    return counter >= n
}