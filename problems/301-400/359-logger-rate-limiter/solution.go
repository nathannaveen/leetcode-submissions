type Logger struct {
    m map[string] int
}


/** Initialize your data structure here. */
func Constructor() Logger {
    return Logger{map[string] int{}}
}


/** Returns true if the message should be printed in the given timestamp, otherwise returns false.
        If this method returns false, the message will not be printed.
        The timestamp is in seconds granularity. */
func (this *Logger) ShouldPrintMessage(timestamp int, message string) bool {
    _, contains := this.m[message]
    if !contains || timestamp - this.m[message] >= 10 {
        this.m[message] = timestamp
        return true
    }
    
    return false
}


/**
 * Your Logger object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.ShouldPrintMessage(timestamp,message);
 */