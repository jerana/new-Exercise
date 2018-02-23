package main

import "fmt"

//Implement a basic calculator to evaluate a simple expression string.

//The expression string contains only non-negative integers, +, -, *, /
//operator//s and empty spaces . The integer division should truncate toward
//zero.

//You may assume that the given expression is always valid.

//Some examples:
//"3+2*2" = 7
//" 3/2 " = 1
//" 3+5 / 2 " = 5

func main() {
	fmt.Println("ex1 :", calculate("3+2*2"))
	fmt.Println("ex2 :", calculate("3/2"))
	fmt.Println("ex3 :", calculate("3+5/2"))
}

type stack []interface{}

func (s *stack) Push(n interface{}) {
	*s = append(*s, n)
}
func (s *stack) Pop() interface{} {
	l := len(*s) - 1
	r := (*s)[l]
	*s = (*s)[0:l]
	return r
}
func (s *stack) Peek() interface{} {
	return (*s)[len(*s)-1]
}
func (s *stack) isEmpty() bool {
	if len(*s) <= 0 {
		return true
	}
	return false
}

func calculate(s string) int {
	var numStack stack
	var opStack stack
	for i := 0; i < len(s); {
		c := s[i]
		if isDigit(c) { //if it's digit then calculated full digit
			j := i
			num := 0
			for j < len(s) && isDigit(s[j]) {
				num += num*10 + int(s[j]-'0')
				j++
			}
			i = j
			numStack.Push(num)
		} else if c == '+' || c == '-' || c == '/' || c == '*' {
			for opStack.isEmpty() != true &&
				precendense(c, opStack.Peek().(byte)) != true {
				op := opStack.Pop().(byte)
				oper1 := numStack.Pop().(int)
				oper2 := numStack.Pop().(int)
				//To keep Operator precendes , pop operator till operator precende match
				numStack.Push(computation(op, oper1, oper2))

			}
			opStack.Push(c)
			i++
		} else {
			i++
		}

	}
	for opStack.isEmpty() != true {
		op := opStack.Pop().(byte)
		oper1 := numStack.Pop().(int)
		oper2 := numStack.Pop().(int)
		numStack.Push(computation(op, oper1, oper2))
	}
	return (numStack.Pop().(int))
}
func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

var hMap map[byte]int = map[byte]int{
	'-': 1,
	'+': 2,
	'/': 3,
	'*': 4,
}

func precendense(cur, pre byte) bool {
	if hMap[cur] > hMap[pre] {
		return true
	}
	return false
}
func computation(op interface{}, op1, op2 int) int {
	switch op.(byte) {
	case '+':
		return op1 + op2
	case '-':
		return op2 - op1
	case '*':
		return op1 * op2
	case '/':
		return op2 / op1

	}
	return 0
}
