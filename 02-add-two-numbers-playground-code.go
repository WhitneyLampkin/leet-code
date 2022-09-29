// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	fmt.Println("Inside addTwoNumbers...")

	// Create variables to hold the current ListNodes, the result linked list and an indicator the carrying value
	l1Node := l1
	l2Node := l2
	sum := 0
	carry := 0
	var summedNode ListNode
	var resultHeadNode ListNode
	currentNode := resultHeadNode

	// Iterate over both linked lists using the first nodes provided
	for {
		// Check both of the first nodes for values
		// Add if both exist (carry if necessary) or keep existing single value
		// Insert each summation to the result linked list by adding a pointer to the new value
		if l1Node.Val != 0 && l2Node.Val != 0 {
			sum = l1Node.Val + l2Node.Val + carry
		} else if l1Node.Val != 0 && l2Node.Val == 0 {
			sum = l1Node.Val + carry
		} else if l1Node.Val == 0 && l2Node.Val != 0 {
			sum = l2Node.Val + carry
		} else if carry > 0 {
			sum = carry
		} else {
			// Break if we've reached the end
			break
		}

		// Reset carry value after each iteration
		carry = 0

		// Reset carry for next iteration
		if sum >= 10 {
			sum = sum - 10
			carry = 1
		}

		// Set list node to add to list
		summedNode = ListNode{Val: sum, Next: nil}

		// Add summed node to result list
		if &resultHeadNode.Val == nil {
			resultHeadNode = summedNode
			currentNode = resultHeadNode
		} else {
			currentNode.Next = &summedNode
			currentNode = summedNode
		}

		// Move to next list nodes
		l1Node = l1Node.Next
		l2Node = l2Node.Next
	}

	// Return the result linked list
	fmt.Println(l1Node, l2Node)
	return &resultHeadNode
}

func main() {
	fmt.Println("Hello, 世界")

	nodeA := &ListNode{Val: 3, Next: nil}
	nodeB := &ListNode{Val: 4, Next: nodeA}
	nodeC := &ListNode{Val: 2, Next: nodeB}

	nodeD := &ListNode{Val: 4, Next: nil}
	nodeE := &ListNode{Val: 6, Next: nodeD}
	//nodeF := &ListNode{Val: 5, Next: nodeE}

	fmt.Println(addTwoNumbers(nodeC, nodeE))
}
