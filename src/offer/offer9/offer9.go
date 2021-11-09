package main

import "fmt"

func main() {
	queue := CQueue{}
	queue.AppendTail(1)
	queue.AppendTail(2)
	queue.AppendTail(3)
	queue.AppendTail(4)
	fmt.Println("pop =",queue.DeleteHead())
	fmt.Println("pop =",queue.DeleteHead())
	fmt.Println("pop =",queue.DeleteHead())
	fmt.Println("pop =",queue.DeleteHead())
	fmt.Println("pop =",queue.DeleteHead())
}

type CQueue struct {
	inputStack		[]int
	outputStack		[]int
}


func Constructor() CQueue {
	return CQueue{
		inputStack: []int{},
		outputStack: []int{},
	}
}


func (this *CQueue) AppendTail(value int)  {
	//inputStack入栈（从数组尾部添加元素）
	this.inputStack = append(this.inputStack,value)
}


func (this *CQueue) DeleteHead() int {
	if len(this.outputStack) == 0 {
		if len(this.inputStack) == 0{
			//如果输入和输出栈都为空
			return -1
		} else {
			//否则迭代输入栈中所有元素到输出栈中
			for len(this.inputStack) != 0 {
				//inputStack出栈
				inputStackPop := this.inputStack[len(this.inputStack)-1]	//获取输入栈栈顶元素
				this.inputStack = this.inputStack[:len(this.inputStack)-1]	//缩减栈的大小
				//inputStack完成出栈

				//outputStack入栈
				this.outputStack = append(this.outputStack,inputStackPop)
			}
		}
	}
	//获取输出栈栈顶元素
	outputStackPop := this.outputStack[len(this.outputStack)-1]
	//输出栈缩减
	this.outputStack = this.outputStack[:len(this.outputStack)-1]
	return outputStackPop
}
