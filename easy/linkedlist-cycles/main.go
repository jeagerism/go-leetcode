// Given head, the head of a linked list, determine if the linked list has a cycle in it.
// There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.
// Return true if there is a cycle in the linked list. Otherwise, return false.

// Example 1:
// Input: head = [3,2,0,-4], pos = 1
// Output: true
// Explanation: There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).

// Example 2:
// Input: head = [1,2], pos = 0
// Output: true
// Explanation: There is a cycle in the linked list, where the tail connects to the 0th node.

// Example 3:
// Input: head = [1], pos = -1
// Output: false
// Explanation: There is no cycle in the linked list.

package main

import "fmt"

type list_node struct {
	value int
	next  *list_node
}

func hasCycle(head *list_node) bool {
	if head == nil || head.next == nil {
		return false
	}

	fast := head
	slow := head

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next

		if fast == slow {
			return true
		}
	}
	return false
}

func createLinkedListWithCycle(values []int, pos int) *list_node {
	if len(values) == 0 {
		return nil
	}

	dummy := &list_node{}
	current := dummy
	var cycleNode *list_node

	for i, val := range values {
		newNode := &list_node{value: val}
		current.next = newNode
		current = newNode
		if i == pos {
			cycleNode = newNode
		}
	}

	if pos != -1 {
		current.next = cycleNode
	}

	return dummy.next
}

func main() {
	head := createLinkedListWithCycle([]int{3, 2, 0, -4}, 1)
	fmt.Println("Has cycle:", hasCycle(head)) // Output: true

	head2 := createLinkedListWithCycle([]int{1, 2}, 0)
	fmt.Println("Has cycle:", hasCycle(head2)) // Output: true

	head3 := createLinkedListWithCycle([]int{1}, -1)
	fmt.Println("Has cycle:", hasCycle(head3)) // Output: false
}
