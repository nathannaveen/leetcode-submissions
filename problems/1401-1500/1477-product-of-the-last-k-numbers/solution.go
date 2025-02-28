type key struct {
    i int
    p int
}

type ProductOfNumbers struct {
    arr []int
    x map[int] key
}


func Constructor() ProductOfNumbers {
    return ProductOfNumbers{[]int{}, make(map[int] key)}
}


func (this *ProductOfNumbers) Add(num int)  {
    this.arr = append(this.arr, num)
}


func (this *ProductOfNumbers) GetProduct(k int) int {
    n := len(this.arr) - k
    if _, ok := this.x[n]; !ok {
        this.x[n] = key{len(this.arr) - k, 1}
    }

    newP := this.x[n].p

    for i := this.x[n].i; i < len(this.arr); i++ {
        newP *= this.arr[i]
    }
    this.x[n] = key{len(this.arr), newP}
    return this.x[n].p
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
