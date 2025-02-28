type key struct {
    start, end string
}

type key2 struct {
    start string
    t int
}

type key3 struct {
    sum int
    n int
}

type UndergroundSystem struct {
    m map[int] key2
    averageTime map[key] key3
}


func Constructor() UndergroundSystem {
    return UndergroundSystem{map[int] key2 {}, map[key] key3 {}}
}


func (this *UndergroundSystem) CheckIn(id int, stationName string, t int)  {
    this.m[id] = key2{stationName, t}
}


func (this *UndergroundSystem) CheckOut(id int, stationName string, t int)  {
    s := this.averageTime[key{this.m[id].start, stationName}].sum + (t - this.m[id].t)
    n := this.averageTime[key{this.m[id].start, stationName}].n + 1
    this.averageTime[key{this.m[id].start, stationName}] = key3{s, n}
}


func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
    return float64(this.averageTime[key{startStation, endStation}].sum) / float64(this.averageTime[key{startStation, endStation}].n)
}


/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */