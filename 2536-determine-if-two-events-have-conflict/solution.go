func haveConflict(event1 []string, event2 []string) bool {
    return event1[0] <= event2[1] && event2[0] <= event1[1]
}