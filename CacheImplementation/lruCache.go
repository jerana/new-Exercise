package main

import (
	"errors"
	"fmt"
)

func main() {
	var cache_sz uint
	fmt.Println("Enter Size of Cache:")
	fmt.Scan(&cache_sz)
	obj := constructor(cache_sz)
	fmt.Println("Cache size:", cache_sz, ":get init")
	obj.put(1, "one")
	obj.put(2, "two")
	v, err := obj.get(5)
	fmt.Println(v, err)
	obj.put(1, "one:one")
	obj.put(3, "three")
	obj.put(4, "four")
	obj.put(2, "two:two")
	v, err = obj.get(4)
	fmt.Println(v, err)
}

/*
Hashtable + Double linked list
The problem can be solved with a hashtable that keeps track of the keys and its values in the double linked list. One interesting property about double linked list is that the node can remove itself without other reference. In addition, it takes constant time to add and remove nodes from the head or tail.

One particularity about the double linked list that I implemented is that I create a pseudo head and tail to mark the boundary, so that we don't need to check the NULL node during the update. This makes the code more concise and clean, and also it is good for the performance as well.
*/

//Double llink list
type llist struct {
	key  int
	s    string
	prev *llist
	next *llist
}

type lruCache struct {
	capacity uint
	count    uint
	headQ    *llist
	tailQ    *llist
	hashTbl  map[int]*llist
}

func constructor(capacity uint) *lruCache {
	c := new(lruCache)
	c.capacity = capacity
	c.headQ = new(llist) //Dummy head and tail to mark boundry
	c.tailQ = new(llist)
	c.headQ.next = c.tailQ
	c.tailQ.prev = c.headQ
	c.hashTbl = make(map[int]*llist, 0) // init HashTable
	return c
}

//Write method to update this cache table
//Add node always put node into head of double llist
func (this *lruCache) addNode(n *llist) {
	//First update nodes ptr
	n.prev = this.headQ
	n.next = this.headQ.next
	//Update Queue head ptrs, make Q head previous node  link to new node
	this.headQ.next.prev = n
	this.headQ.next = n
}

//Remove node from Double list
func (this *lruCache) removeNode(n *llist) {
	//record nodes previous and next nodes
	pre := n.prev
	nxt := n.next
	//Update previos nad next nodes and let Garbage collectore to free node
	pre.next = nxt
	nxt.prev = pre
}

//Case: whenever node is touched , should moved to head of Q list
func (this *lruCache) moveTohead(n *llist) {
	//first delete the nodes from llist
	this.removeNode(n)
	//Add the node into Head of list
	this.addNode(n)
}

//Case: Lru the tail node
func (this *lruCache) popTail() *llist {
	//Get Tail Node
	rm := this.tailQ.prev
	//Call remove this node
	this.removeNode(rm)
	return rm
}

//Get function :
func (this *lruCache) get(key int) (string, error) {
	n, ok := this.hashTbl[key]
	if !ok { //Key doesn't exist
		return "", errors.New("Key doesn't exist")
	}
	//Move this node to head since it got touched
	this.moveTohead(n)
	return n.s, nil
}
func (this *lruCache) delCache(key int) {
	this.count--
	delete(this.hashTbl, key)
}
func (this *lruCache) addCache(key int, n *llist) {
	this.count++
	this.hashTbl[key] = n
}

//Set function:
func (this *lruCache) put(key int, str string) {
	n, ok := this.hashTbl[key]
	if ok {
		n.s = str
		this.moveTohead(n)
	} else { //New key
		n = new(llist)
		n.key = key
		n.s = str
		this.addNode(n)
		this.addCache(key, n)
		if this.count > this.capacity { //Does require LRU
			r := this.popTail()
			this.delCache(r.key) //will remove node from cache
		}
	}
}
