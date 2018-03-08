package main

import "fmt"

func main() {
	fmt.Println("Request Parenthsis 2  ")
	fmt.Println(findPossibleParenthis(2))
}
func findPossibleParenthis(n int) [][]byte {
	cur := make([]byte, 2*n)
	result := make([][]byte, 0)
	buildValidParenthis(cur, 0, &result)
	return result
}
func buildValidParenthis(cur []byte, po int, res *[][]byte) {
	if po == len(cur) {
		if isValidParenthsisString(cur) {
			//if isValid(cur) {
			t := make([]byte, len(cur))
			copy(t, cur)
			*res = append(*res, t)
			fmt.Println("Add to result", t, res)
		}
	} else {
		cur[po] = '('
		buildValidParenthis(cur, po+1, res)
		cur[po] = ')'
		buildValidParenthis(cur, po+1, res)
	}
}

type stack []interface{}

func (s *stack) Push(c interface{}) {
	*s = append(*s, c)
}
func (s *stack) Pop() interface{} {
	c := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return c
}
func (s *stack) isEmpty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}
func (s *stack) Peek() interface{} {
	return (*s)[len(*s)-1]
}
func isValid(s []byte) bool {
	balance := 0
	for _, c := range s {
		if c == '(' {
			balance += 1
		} else {
			balance -= 1
		}
		if balance < 0 {
			return false
		}
	}
	if balance == 0 {
		return true
	}
	return false
}

func isValidParenthsisString(s []byte) bool {
	var st stack
	for _, c := range s {
		//fmt.Println("%v %v %v", c, ')', '(')
		if c == ')' {
			if st.isEmpty() != true && st.Peek().(byte) == '(' {
				st.Pop()
			} else {
				st.Push(c)
			}
		} else {
			st.Push(c)
		}
	}
	//fmt.Println("given str:%v, and stack:%v", s, st)
	if st.isEmpty() {
		return true
	}
	return false

}
