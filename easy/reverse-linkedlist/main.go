// Input: head = [1,2,3,4,5]
// Output: [5,4,3,2,1]

package main

func main() {
	head := linklist([]int{1, 2, 3, 4, 5})
	printList(head)

	// reverse แล้วเก็บผลลัพธ์
	reversed := reverseList(head)
	printList(reversed)
}

type list_node struct {
	val  int
	next *list_node
}

func linklist(value []int) *list_node {
	head := &list_node{}
	current := head
	for _, val := range value {
		newNode := &list_node{val: val}
		current.next = newNode
		current = newNode
	}
	return head.next
}

func printList(head *list_node) {
	for head != nil {
		print(head.val)
		if head.next != nil {
			print(" -> ")
		}
		head = head.next
	}
	println()
}

func reverseList(head *list_node) *list_node {
	current := head
	var prev, next *list_node

	for current != nil {
		next = current.next
		current.next = prev

		prev = current
		current = next
	}

	return prev
}
