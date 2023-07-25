/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

 

Example 1:


Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
 

Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // Create variables to hold the current ListNodes, the result linked list and an indicator the carrying value
	l1Node := l1
	l2Node := l2
	sum := 0
	carry := 0
	var summedNode *ListNode
	var resultHeadNode *ListNode
	//pointer := ListNode{Val: 0, Next: nil}

	// Iterate over both linked lists using the first nodes provided
	for {
		// Check both of the first nodes for values
		// Add if both exist (carry if necessary) or keep existing single value
		// Insert each summation to the result linked list by adding a pointer to the new value
		if l1Node != nil && l2Node != nil {
			sum = l1Node.Val + l2Node.Val + carry
		} else if l1Node != nil && l2Node == nil {
			sum = l1Node.Val + carry
		} else if l1Node == nil && l2Node != nil {
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
		summedNode = &ListNode{Val: sum, Next: nil}
		fmt.Println(summedNode.Val)

		if resultHeadNode == nil {
		    resultHeadNode = summedNode
		} else {
		    pointer := resultHeadNode
		    for pointer.Next != nil {
			pointer = pointer.Next
		    }
		    pointer.Next = summedNode
		}

			// Move to next list nodes
		if l1Node != nil {
		    l1Node = l1Node.Next
		}
			if l2Node != nil {
		    l2Node = l2Node.Next
		}
	}

	// Return the result linked list
	return resultHeadNode
}
