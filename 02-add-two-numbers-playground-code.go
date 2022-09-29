// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
	len int
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	fmt.Println("Inside addTwoNumbers...")

	l1ListNode := l1
	l2ListNode := l2
	resultListNode := &ListNode{Val: 0, Next: nil}
	var currentNode *ListNode

	firstIteration := true
	carry := 0

	for l1ListNode != nil || l2ListNode != nil {
		if l1ListNode != nil && l2ListNode != nil {
			summedVal := l1ListNode.Val + l2ListNode.Val

			if summedVal >= 10 {
				carry = 1
				summedVal = summedVal - 10
			} else {
				carry = 0
			}

			if firstIteration {
				resultListNode := &ListNode{Val: summedVal, Next: nil}
				currentNode = resultListNode
			} else if resultListNode.Next == nil {
				resultListNode.Next = &ListNode{Val: summedVal, Next: nil}
				currentNode = resultListNode.Next
			} else {
				currentNode.Next = &ListNode{Val: summedVal, Next: nil}
			}
		} else if l1ListNode != nil && l2ListNode == nil {

		} else if l1ListNode == nil && l2ListNode != nil {

		}

		if l1ListNode.Next != nil {
			l1ListNode = l1ListNode.Next
		}
		if l2ListNode.Next != nil {
			l2ListNode = l2ListNode.Next
		}
		
		firstIteration = false
	}

	fmt.Println(carry)
	fmt.Println(resultListNode)
	return resultListNode
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

