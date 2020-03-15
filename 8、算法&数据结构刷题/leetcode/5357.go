type CustomStack struct {
    ValList []int
    maxSize int
    curSize int
}


func Constructor(maxSize int) CustomStack {
    newStack := CustomStack{}
    newStack.maxSize = maxSize
    return newStack
}


func (this *CustomStack) Push(x int)  {
    if this.curSize < this.maxSize {
        this.ValList = append(this.ValList, x)
        this.curSize ++
    }
}


func (this *CustomStack) Pop() int {
    if this.curSize > 0 {
        vv := this.ValList[this.curSize-1]
        this.ValList = this.ValList[0:this.curSize-1]
        this.curSize --
        return vv
    }
    return -1
}


func (this *CustomStack) Increment(k int, val int)  {
    for i:=0; i<k && i<this.curSize; i++ {
        this.ValList[i] += val
    }
}


/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */