package main

import "math"

func main() {

}

type StackNode struct {
	Val  int
	Min  int
	Next *StackNode
}

type MinStack struct {
	head *StackNode
}

// Constructor /** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{head: nil}
}

func (s *MinStack) Push(x int) {
	if s.head == nil {
		s.head = &StackNode{
			Val: x,
			Min: x,
		}
		return
	}
	s.head.Next = &StackNode{
		Val:  x,
		Min:  Min(x, s.Min()),
		Next: s.head,
	}
}

func (s *MinStack) Pop() {
	if s.head == nil {
		return
	}
	s.head = s.head.Next
}

func (s *MinStack) Top() int {
	if s.head == nil {
		return math.MaxInt
	}

	return s.head.Val
}

func (s *MinStack) Min() int {
	if s.head == nil {
		return math.MaxInt
	}

	return s.head.Min
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Min();
 */
