func maxHeightOfTriangle(red int, blue int) int {
    cnt := helper(red, blue)
    cnt2 := helper(blue, red)

    return max(cnt, cnt2)
}

func helper(red, blue int) int {
    x := 1
    z := true
    cnt := 0

    for red >= 0 && blue >= 0 {
        if z {
            if blue >= x {
                cnt++
                blue -= x
            } else {
                break
            }
        } else {
            if red >= x {
                cnt++
                red -= x
            } else {
                break
            }
        }
        x++
        z = !z
    }

    return cnt
}
