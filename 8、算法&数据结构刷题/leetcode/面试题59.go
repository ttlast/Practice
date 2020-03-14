type MaxQueue struct {
	valList 	[]int
	maxRank 	[]int
	popPos		int
	rankTop 	int
	rankPos 	int
	rankSize 	int
	size		int
}


func Constructor() MaxQueue {
	return MaxQueue{}
}


func (this *MaxQueue) Max_value() int {
	if this.popPos >= this.size {
		return -1
	}

	return this.valList[this.maxRank[this.rankTop]]
}


func (this *MaxQueue) Push_back(value int)  {
	this.valList = append(this.valList, value)

	for this.rankTop < this.rankPos {
		if (this.valList[this.maxRank[this.rankPos-1]] > value) {
            break
        }
		this.rankPos --
	}

	if (this.rankPos >= this.rankSize) {
		this.maxRank = append(this.maxRank, this.size)
		this.rankSize ++
	} else {
		this.maxRank[this.rankPos] = this.size
	}

	this.rankPos ++

	this.size ++
}


func (this *MaxQueue) Pop_front() int {
	if this.popPos >= this.size {
		return -1
	}

	retVal := this.valList[this.popPos]

	this.popPos ++
	if this.popPos > this.maxRank[this.rankTop] {
		this.rankTop ++
	}

	return retVal
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */