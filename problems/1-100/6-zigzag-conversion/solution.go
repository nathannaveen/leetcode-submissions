func convert(s string, numRows int) string {
    
    if numRows == 1 { return s }
    
    arr := make([]string, numRows)
    minus := false
    counter := 0
    
    for i := range s {
        arr[counter] += string(s[i])
        if !minus {
            counter++
            if counter == len(arr) {
                minus = true
                counter -= 2
            }
        } else {
            counter--
            if counter == -1 {
                minus = false
                counter += 2
            }
        }
    }
    return strings.Join(arr, "")
}


// p i n
// a l s i g
// y a h r
// p i