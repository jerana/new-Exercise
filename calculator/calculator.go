package main

//Problem : Given a string of format '2+3*2-1', calculate and return  the
//result. No parenthesis in the input, just integers and + - * / operators.
//Operator precedence has to be considered. Linear time complexity and minimal
//data structure use is preferred.
import "fmt"

type stack []interface{}

func (s *stack) Push(n interface{}) {
	*s = append(*s, n)
}
func (s *stack) Pop() interface{} {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
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
func operandAndOperatorResult(op1, op2 int, op byte) int {
	if op == '+' {
		return op1 + op2
	} else if op == '-' {
		return op1 - op2
	} else if op == '*' {
		return op1 * op2
	}
	return op1 / op2
}
func getPrecendence(op byte) int {
	if op == '-' {
		return 1
	} else if op == '+' {
		return 2
	} else if op == '/' {
		return 3
	}
	return 4
}

func main() {
	s := "2+3*2-1"
	fmt.Println(s)
	fmt.Println(calculator(s))
}
func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}

func calculator(s string) int {
	var operand stack
	var operator stack
	for i := 0; i < len(s); i++ {
		if isDigit(s[i]) {
			num := 0
			j := 0
			for j = i; j < len(s) && isDigit(s[j]); j++ {
				num += num*10 + int(s[j]-'0')
			}
			i = j - 1
			fmt.Println("Push Operand :", num)
			operand.Push(num)
		} else {
			fmt.Println("Got Operator  ")
			//Check current operator precedence with previous
			//operarator, pop old operator and its operand from
			//other stack and push result back to operand stack if
			//new operator has lower precendce than old operator
			for {
				if operator.isEmpty() != true {
					old := operator.Peek()
					fmt.Println("old Operator:", old)
					fmt.Println("new Operator", s[i])
					if getPrecendence(old.(byte)) > getPrecendence(byte(s[i])) {
						fmt.Println("New operator precedenc is less than old")
						op1 := operand.Pop()
						op2 := operand.Pop()
						op := operator.Pop()
						r := operandAndOperatorResult(op2.(int), op1.(int), op.(byte))
						fmt.Println("push the new calculated res::", r)
						operand.Push(r)
					} else {
						fmt.Println("Break from operator loop:")
						break
					}
				} else {
					fmt.Println("Operator stack is empty")
					break
				}
			}
			fmt.Println("Push Operator into stack", s[i])
			operator.Push(s[i])
			fmt.Println("Operator stack", operator)

		}

	}
	//Walk to operator Stack and calculate final result
	for operator.isEmpty() != true {
		op := operator.Pop()
		op1 := operand.Pop()
		op2 := operand.Pop()
		r := operandAndOperatorResult(op2.(int), op1.(int), op.(byte))
		operand.Push(r)
	}
	return (operand.Pop()).(int)
}

//"((2+5) * 7 + 3-1)"
func calculator_part_2(s string) int {
	sing := '+'
	var calStack stack
	stack.Push(0)
	for i := 0; i < len(s); i++ {
		num = 0
		if s[i] == '(' {
			k = i
			j = i + 1
			if s[j] == '(' {
				k = j
				j++
			} else if s[j] == ')' {
				k -= 1
				num := calculator_part_2(s[k:j])
			}

		}
		if isDigit(s[i]) {
			for j = i; j < len(s) && isDigit(s[j]); j++ {
				num += num*10 + s[j] - '0'
			}
			result := op1

		}
	}

}
func calculatorString(op1, op2 byte, sign byte) int {
	if sign == '+' {
		return op1 + op2
	} else if sign == '_' {
		return op1 * io1

	} else if sign  == ')' {
		return  0p1*op2
	}
	return op1/op2
}
