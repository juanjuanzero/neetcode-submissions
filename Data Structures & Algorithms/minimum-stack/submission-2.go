

type MinStack struct {
	Min  []int
	Vals []int
}

func NewMinStack() MinStack {
	return MinStack{}
}

func Constructor() MinStack {
	return NewMinStack()
}

func (this *MinStack) Push(val int) {
	// if the current min is greater than or equal to the incoming value
	// add it to the stack
	if len(this.Min) == 0 && len(this.Vals) == 0 {
		this.Min = []int{val}
		this.Vals = []int{val}
		return
	}

	currMin := this.Min[len(this.Min)-1]
	if currMin >= val {
		// add it to the min stack
		this.Min = append(this.Min, val)
	}

	this.Vals = append(this.Vals, val)
}

func (this *MinStack) Pop() {
	lastVal := this.Vals[len(this.Vals)-1]
	// is lastVal greater than the value of the top of Min stack?
	lastMin := this.Min[len(this.Min)-1]
	if lastVal == lastMin {
		// also pop out of last min
		if len(this.Min) > 1 {
			this.Min = this.Min[:len(this.Min)-1]
		} else {
			this.Min = []int{}
		}
	}
	// finally pop out of the stack
	if len(this.Vals) > 1 {
		this.Vals = this.Vals[:len(this.Vals)-1]
	} else {
		this.Vals = []int{}
	}

}

// returns the top of the min stack
func (this *MinStack) Top() int {
	return this.Vals[len(this.Vals)-1]
}

func (this *MinStack) GetMin() int {
	return this.Min[len(this.Min)-1]
}
