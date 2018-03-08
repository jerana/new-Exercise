package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

//Stack Implementation Using Slice in Go
type Stack struct {
	lock sync.Mutex
	s    []interface{}
}

func NewStack() *Stack {
	return &Stack{syc.Mutext{}, make([]interface{}, 0)}
}

func (s *Stack) Push(v interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.s = append(s.s, v)
}
func (s *Stack) Pop() (interface{}, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.s) == 0 {
		return -1, errors.New("Empty Stack")
	}
	res := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return res, nil
}
func (s *Stack) IsEmpty() bool {
	s.lock.Lock()
	defer s.lock.Unlock()
	if len(s.s) == 0 {
		return true
	}
	return false
}

type tree struct {
	key   int
	left  *tree
	right *tree
}

func InorderWalk(root *tree) []*tree {
	t := make([]*tree, 0)
	s := NewStack()
	n := root
	s.Push(n)
	for {
		if s.IsEmpty() {
			return t
		}
		if n.left != nil {
			n := n.left
			s.Push(n)
		} else {
			n = s.Pop()
			t = append(t, n)
			if n.right != nil {
				n := n.right
				s.push(n)
			} else {
				n = s.Pop()
			}
		}
	}
}
