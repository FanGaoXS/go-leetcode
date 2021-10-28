package main

import "fmt"

func main() {
	queue := &MyQueue{}
	fmt.Println("empty =",queue.Empty())
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	queue.Push(4)
	fmt.Println("empty =",queue.Empty())
	fmt.Println("peek =", queue.Peek())
	fmt.Println("pop =", queue.Pop())
	fmt.Println("pop =", queue.Pop())
	fmt.Println("pop =", queue.Pop())
	fmt.Println("pop =", queue.Pop())
	fmt.Println("empty =",queue.Empty())
}

type MyQueue struct {
	inputStack			[]int
	outputStack			[]int
}


func Constructor() MyQueue {
	return MyQueue{}
}


func (this *MyQueue) Push(x int)  {
	this.inputStack = append(this.inputStack,x)
}


func (this *MyQueue) Pop() int {
	this.ensureOutputStack()
	pop := this.outputStack[len(this.outputStack)-1]
	this.outputStack = this.outputStack[:len(this.outputStack)-1]
	return pop
}


func (this *MyQueue) Peek() int {
	this.ensureOutputStack()
	return this.outputStack[len(this.outputStack)-1]
}


func (this *MyQueue) Empty() bool {
	return len(this.inputStack) == 0  && len(this.outputStack) ==0
}

func (this *MyQueue) ensureOutputStack()  {
	if len(this.outputStack) == 0 {
		for len(this.inputStack) != 0 {
			inputLength := len(this.inputStack)
			//inputStack出栈
			pop := this.inputStack[inputLength-1]
			this.inputStack = this.inputStack[:inputLength-1]
			//outputStack入栈
			this.outputStack = append(this.outputStack,pop)
		}
	}
}



