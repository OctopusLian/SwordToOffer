type MaxQueue struct {
	q1  []int
	max []int
}

func Constructor() MaxQueue {
	return MaxQueue{
		make([]int, 0),
		make([]int, 0),
	}
}

func (this *MaxQueue) Max_value() int {
	if len(this.max) == 0 {
		return -1
	}
	return this.max[0]
}

func (this *MaxQueue) Push_back(value int) {
	this.q1 = append(this.q1, value) //将value入q1队列
	for len(this.max) != 0 && value > this.max[len(this.max)-1] {
		this.max = this.max[:len(this.max)-1]
	}
	this.max = append(this.max, value)
}

func (this *MaxQueue) Pop_front() int {
	n := -1
	if len(this.q1) != 0 {
		n = this.q1[0]
		this.q1 = this.q1[1:]
		if this.max[0] == n {
			this.max = this.max[1:]
		}
	}
	return n
}

/**
 * Your MaxQueue object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Max_value();
 * obj.Push_back(value);
 * param_3 := obj.Pop_front();
 */