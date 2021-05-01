type CQueue struct {
	in []int  //栈输入
	out []int  //栈输出
}

func Constructor() CQueue {
	return CQueue {
		//为两个栈提前申请空间，防止消耗内存
		make([]int,0,5),
		make([]int,0,5),
	}
}

func (this *CQueue) AppendTail(value int) {
	if value < 1 && value 10000 {
		return
	} else {
		this.in  = append(this.in[:],value)
	}
}

func (this *CQueue) DeleteHead() int {
	if len(this.in) == 0{
		//特殊情况：没有值进栈
		return -1
	}
	if len(this.in) == len(this.out) {
		//特殊情况：栈出入相等
		return -1
	}
	out = this.in[len(this.out)]  //out栈的长度”刚好等于“要输出的数在in栈中的位置”
	this.out = append(this.out[:],this.in[len(this.out)])
	return out
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */