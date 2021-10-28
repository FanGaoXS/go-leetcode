package main

import (
	"fmt"
	"math"
)

func main() {
	minStack := new(MinStack)
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	fmt.Println("getMin() =",minStack.GetMin())
	minStack.Pop()
	minStack.Top()
	fmt.Println("getMin() =",minStack.GetMin())
}

type MinStack struct {
	//主栈
	stack			[]int
	//辅助栈
	minStack 		[]int
}


func Constructor() MinStack {
	return MinStack{
		minStack: []int{math.MaxInt32},
	}
}


func (this *MinStack) Push(val int)  {
	this.stack = append(this.stack,val)
	if val < this.minStack[len(this.minStack)-1] {
		this.minStack = append(this.minStack,val)
	} else {
		this.minStack = append(this.minStack,this.minStack[len(this.minStack)-1])
	}
}


func (this *MinStack) Pop()  {
	this.stack = this.stack[:len(this.stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
