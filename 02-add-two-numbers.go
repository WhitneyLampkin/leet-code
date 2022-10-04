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
