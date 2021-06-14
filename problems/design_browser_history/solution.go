type BrowserHistory struct {
    history []string
    position int
}


func Constructor(homepage string) BrowserHistory {
    b := BrowserHistory{}
    b.history, b.position = append(b.history, homepage), 1
    return b
}


func (this *BrowserHistory) Visit(url string)  {
    if len(this.history) > this.position {
        this.history = this.history[:this.position + 1]
    }
    this.history = append(this.history, url)
    this.position = len(this.history) - 1
}


func (this *BrowserHistory) Back(steps int) string {
    this.position -= int(math.Min(float64(steps), float64(this.position)))
    return this.history[this.position]
}


func (this *BrowserHistory) Forward(steps int) string {
    this.position += int(math.Min(float64(steps), float64(len(this.history) - this.position - 1)))
    return this.history[this.position]
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */