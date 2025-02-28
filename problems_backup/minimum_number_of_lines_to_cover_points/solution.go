type line struct {
    slope float64
    intercept float64 // y intercept
}

type point struct {
    x, y int
}

func minimumLines(points [][]int) int {
    used := make(map[point] bool)
    numLines := 0
    
    for len(used) != len(points) {
        startPoint := point{ 0, 0 }
        a := []point{}
        
        for _, p1 := range points {
            m := make(map[line] []point) // If the point has the same slope and y intercept we can add it to m
            if used[point{ p1[0], p1[1] }] {
                // fmt.Println(p1, "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX")
                continue
            }
            for _, p2 := range points {
                x1, y1, x2, y2 := p1[0], p1[1], p2[0], p2[1]
                
                if x1 == x2 && y1 == y2 {
                //     startPoint = p1
                    continue
                }
                
                temp := point{ x2, y2 }
                slope := float64(y2 - y1) / float64(x2 - x1)
                intercept := float64(y2) - slope * float64(x2)
                
                if x1 == x2 {
                    slope = float64(100000)
                    intercept = float64(100000)
                }
                
                // fmt.Println("Slope:", slope, "Intercept:", intercept, "p2:", p2, "p1:", p1)

                k := line{ slope, intercept }

                if !used[temp] {
                    m[k] = append(m[k], temp)
                }
            }
            arr := []point{}
            // key := line{ float64(0), float64(0) }

            for _, b := range m {
                if len(b) >= len(arr) {
                    arr = b
                    // key = a
                }
            }
            if len(arr) >= len(a) {
                a = arr
                startPoint = point{p1[0], p1[1]}
            }
//             fmt.Println("map {")
            
//             for a, b := range m {
//                 fmt.Println(a, b)
//             }
            
//             fmt.Println("} ", p1, "\n")
            
        }
        
        for _, p := range a {
            used[p] = true
        }
        
        used[startPoint] = true
        
//         fmt.Println("Used map {")

//         for a, b := range used {
//             fmt.Println(a, b)
//         }

//         fmt.Println("} ")
        
//         fmt.Println("------------------------   a:", a, "\n")
        
        numLines++
    }
    
    return numLines
}

// [[0,-2],[0,5]]