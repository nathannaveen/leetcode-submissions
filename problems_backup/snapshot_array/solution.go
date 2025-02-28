type key struct {
    val int
    hasSet bool
}

type SnapshotArray struct {
    cnt int
    m map[int] map[int] key // index : [snap_id : {val, hasSet}]
    l int
}


func Constructor(length int) SnapshotArray {
    m := map[int] map[int] key {}

    for i := 0; i < length; i++ {
        m[i] = map[int] key {}
    }

    return SnapshotArray{ 0, m, length}
}


func (this *SnapshotArray) Set(index int, val int)  {
    this.m[index][this.cnt] = key{val, true}
}


func (this *SnapshotArray) Snap() int {
    this.cnt++
    return this.cnt - 1
}


func (this *SnapshotArray) Get(index int, snap_id int) int {
    for i := snap_id; i >= 0; i-- {
        if this.m[index][i].hasSet {
            return this.m[index][i].val
        }
    }
    return 0
}


/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */