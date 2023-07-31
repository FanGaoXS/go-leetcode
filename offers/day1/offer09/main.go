package main

func main() {}

type CQueue struct {
	inputStack  []int
	outputStack []int
}

func Constructor() CQueue {
	return CQueue{
		inputStack:  make([]int, 0, 0),
		outputStack: make([]int, 0, 0),
	}
}

func (q *CQueue) AppendTail(value int) {
	q.inputStack = append(q.inputStack, value)
}

func (q *CQueue) DeleteHead() int {
	if len(q.inputStack) == 0 && len(q.outputStack) == 0 {
		return -1
	}

	if len(q.outputStack) == 0 {
		for len(q.inputStack) != 0 {
			// remove the end element of the input stack
			inputTail := q.inputStack[len(q.inputStack)-1]
			q.inputStack = q.inputStack[:len(q.inputStack)-1]

			q.outputStack = append(q.outputStack, inputTail)
		}
	}

	outputTail := q.outputStack[len(q.outputStack)-1]
	q.outputStack = q.outputStack[:len(q.outputStack)-1]
	return outputTail
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
