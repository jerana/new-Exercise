package main

import "fmt"

//:Problem Defination:
//Implement a basic calculator to evaluate a simple expression string.
//The expression string may contain open ( and closing parentheses )
//, the plus + or minus sign -, non-negative integers and empty spaces .
//The expression string contains only non-negative integers, +, -, *, /
// operators , open ( and closing parentheses ) and empty spaces .
// The integer division should truncate toward zero.

func main() {
	fmt.Println("result:", calculate("2+(5*4+3)/3 -2"))
}

type stack []interface{}

func (s *stack) Push(n interface{}) {
	*s = append(*s, n)
}
func (s *stack) Pop() interface{} {
	p := (*s)[len(*s)-1]
	*s = (*s)[0 : len(*s)-1]
	return p
}
func (s *stack) Peek() interface{} {
	return ((*s)[len(*s)-1])
}
func (s *stack) isEmpty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}

func calculate(s string) int {
	var st stack
	sign := byte('+')
	result := 0
	i := 0
	j := 0
	for i < len(s) {
		c := s[i]
		if c == '(' {
			l := 1
			j = i + 1
			//Get the block
			for j < len(s) && l > 0 {
				if s[j] == '(' {
					l += 1
				} else if s[j] == ')' {
					l -= 1
				}
				j++
			}

			blockVal := calculate(s[i+1 : j-1])
			i = j
			if sign == '+' {
				st.Push(blockVal)
			} else if sign == '-' {
				st.Push(-blockVal)
			} else if sign == '*' {
				p := (st.Pop()).(int)
				st.Push(p * blockVal)
			} else if sign == '/' {
				p := (st.Pop()).(int)
				st.Push(p / blockVal)
			}
		} else if isDigit(c) {
			j = i
			digit := 0
			for j < len(s) && isDigit(s[j]) {
				digit = digit*10 + (int)(s[j]-'0')
				j++
			}
			if sign == '+' {
				st.Push(digit)
			} else if sign == '-' {
				st.Push(-digit)
			} else if sign == '*' {
				st.Push((st.Pop()).(int) * digit)
			} else if sign == '/' {
				st.Push((st.Pop()).(int) / digit)
			}
			i = j
		} else if c == '+' || c == '-' || c == '*' || c == '/' {
			sign = c
			i++
		} else {
			i++
		}
	}
	for st.isEmpty() != true {
		result += (st.Pop()).(int)
	}
	return result
}
func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}
