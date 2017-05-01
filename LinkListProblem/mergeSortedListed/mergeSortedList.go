package main

import "fmt"

type llist struct {
	key  int
	next *llist
}

func appendNode(node, tail **llist) {
	(*tail).next = *node
	(*tail) = *node
	(*node) = (*node).next
}
func mergeSortedList(l1, l2 *llist) *llist {
	var dummy, tail *llist
	dummy = new(llist)
	tail = dummy
	for l1 != nil && l2 != nil {
		if l1.key < l2.key {
			appendNode(&l1, &tail)
		} else {
			appendNode(&l2, &tail)
		}
	}
	if l1 == nil {
		tail.next = l2
	} else if l2 == nil {
		tail.next = l1
	}
	return dummy
}
