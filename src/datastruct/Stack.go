package datastruct

type Stack struct {
	size		int
	elements	[]int
}

func Constructor() Stack {
	return Stack{}
}

func (this *Stack) Size() int {
	return this.size
}

func (this *Stack) Push(val int)  {
	this.elements = append(this.elements,val)
	this.size++
}

func (this *Stack) Pop() int {
	last := this.elements[len(this.elements)-1]
	this.elements = this.elements[:len(this.elements)-1]
	this.size--
	return last
}

func (this *Stack) Peek() int {
	return this.elements[len(this.elements)-1]
}

func (this *Stack) isEmpty() bool {
	return this.Size() == 0
}