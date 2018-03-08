package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

//store Tree All all level node as set of array where array set index represent by tree depth level

type treeNode struct {
	key   int
	left  *treeNode
	right *treeNode
}
type qNode struct {
	node *treeNode
	next *qNode
}
type Queue struct {
	qHead *qNode
	qTail *qNode
	len   int
}

func Queue_init() *Queue {
	obj = new(Queue)
	obj.qHead = new(qNode)
	obj.qTail = qHead
	return obj
}
func (q *Queue) Enqueue(n *treeNode) {
	p := new(qNode)
	p.node = n
	q.qTail.next = p
	q.qTail = p
	q.len++
}
func (q *Queue) Dequeue() *treeNode {
	r := qHead.next
	if r == nil {
		return r
	}
	q.qHead.next = r.next
	q.len--
	return r.node
}
func isEmpty() bool {
	if qHead.len == 0 {
		return true
	}
	return false
}

func treeDepthLevelSet(root *treeNode) [][]*treeNode {
	var levelSet [][]*treeNode
	if root == nil {
		return nil
	}
	treeNode = make([][]*treeNode, 1)
	treeNode[0] = make([]*treeNode, 0)
	var levelCount = 0
	var dIndex = 0
	obj := Queue_init()
	obj.Enqueue(root)
	levelCount = obj.len
	for obj.isEmpty() != true {
		n := obj.Dequeue()
		levelSet[dIndex] = append(levelSet[dIndex], n)
		levelCount--
		if n.left != nil {
			obj.Enqueue(n.left)
		}
		if n.right != nil {
			obj.Enqueue(n.right)
		}

		if levelCount == 0 {
			levelCount = obj.len
			t := make([]*treeNode, 0)
			levelSet = append(levelSet, t)
			dIndex++
		} else {
			levelCount--
		}
	}
}
