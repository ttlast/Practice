type stProject struct{
    Cap int
    Pro int
}

type MaxHeap []stProject
func (h MaxHeap) Len() int { return len(h)}
func (h MaxHeap) Less(i,j int) bool { return h[i].Pro > h[j].Pro }
func (h MaxHeap) Swap(i,j int) { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) {
    *h = append(*h, x.(stProject))
}
func (h *MaxHeap) Pop() interface{} {
    x := (*h)[len(*h)-1]
    *h = (*h)[:len(*h)-1]
    return x
}



func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
    nlen := len(Profits)
    projSlice := make([]stProject, nlen)

    for i:=0; i<nlen; i++ {
        projSlice[i].Cap = Capital[i]
        projSlice[i].Pro = Profits[i]
    }
    sort.Slice(projSlice, func(i, j int) bool {
        return projSlice[i].Cap < projSlice[j].Cap
    })

    var h MaxHeap
    heap.Init(&h)

    num := 0
    ret := W
    i := 0
    for true {
        if i >= nlen { break }
        if projSlice[i].Cap > ret {
            // 计算新的资本
            if h.Len() <= 0  || num >= k {
                break
            }
            item := heap.Pop(&h).(stProject)
            ret += item.Pro
            num ++
        } else {
            heap.Push(&h, projSlice[i])
            i ++;
        }
    }
    //return ret
    for true {
       if h.Len() <= 0  || num >= k {
            break
        }
        item := heap.Pop(&h).(stProject)
        ret += item.Pro
        num ++
    }
    return ret

}