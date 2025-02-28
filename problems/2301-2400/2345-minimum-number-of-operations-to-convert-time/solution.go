func convertTime(current string, correct string) int {
    currentTime := int(current[0] - '0') * 600 + int(current[1] - '0') * 60 + int(current[3] - '0') * 10 + int(current[4] - '0')
    correctTime := int(correct[0] - '0') * 600 + int(correct[1] - '0') * 60 + int(correct[3] - '0') * 10 + int(correct[4] - '0')
    
    diff := correctTime - currentTime
    
    sum := diff / 60
    diff %= 60
    sum += diff / 15
    diff %= 15
    sum += diff / 5
    diff %= 5
    sum += diff
    
    return sum
}