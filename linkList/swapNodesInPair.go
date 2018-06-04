package main

/*Problem Statment:
Given a linked list, swap every two adjacent nodes and return its head.

Example:

Given 1->2->3->4, you should return the list as 2->1->4->3.
Note:

Your algorithm should use only constant extra space.
You may not modify the values in the list's nodes, only nodes itself may be changed.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	dummy := new(ListNode)
	var prev, cur, next *ListNode
	prev = dummy
	dummy.Next = head
	cur = head
	for cur != nil && cur.Next != nil {
		next = cur.Next
		tmp := next.Next
		prev.Next = next
		cur.Next = tmp
		next.Next = cur
		prev = cur
		cur = tmp
	}
	return dummy.Next
}
