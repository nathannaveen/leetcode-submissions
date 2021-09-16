func maximumNumber(num string, change []int) string {
    arr := strings.Split(num, "")
    changed := false 
    
    for i := range arr {
        temp, _ := strconv.Atoi(arr[i])
        if change[temp] > temp {
            arr[i] = strconv.Itoa(change[temp])
            changed = true
        } else if changed == true && change[temp] == temp {
            arr[i] = strconv.Itoa(change[temp])
        } else if changed {
            break
        }
    }
    return strings.Join(arr, "")
}