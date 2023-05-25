func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
    tomatoCnt := 0
    for tomatoSlices != cheeseSlices * 2 && tomatoSlices != 0 && cheeseSlices != 0 {
        tomatoCnt++
        tomatoSlices -= 4
        cheeseSlices--
    }

    if cheeseSlices * 2 != tomatoSlices {
        return []int{}
    }

    return []int {tomatoCnt, cheeseSlices}
}