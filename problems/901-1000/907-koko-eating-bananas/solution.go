func minEatingSpeed(piles []int, h int) int {
    low, high := 0, maxPile(piles)
    mid := 0
    
    for low < high {
        hours := 0 // number of hours needed to eat all bananas with k = mid
        mid = low + (high - low) / 2
        
        if low + 1 == high {
            return high
        }
        
        for _, pile := range piles {
            hours += pile / mid
            if pile % mid != 0 {
                hours += 1
            }
        }
        
        if hours > h {
            low = mid
        } else {
            high = mid
        }
    }
    return mid
}

func maxPile(piles []int) (high int) {
    for _, pile := range piles {
        if pile > high {
            high = pile
        }
    }
    return
}