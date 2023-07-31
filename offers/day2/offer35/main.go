package main

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	first := head
	nodes := make(map[*Node]*Node)

	for head != nil {
		nodes[head] = &Node{Val: head.Val}
		head = head.Next
	}

	head = first
	for head != nil {
		nodes[head].Next = nodes[head.Next]
		nodes[head].Random = nodes[head.Random]

		head = head.Next
	}

	return nodes[first]
}
