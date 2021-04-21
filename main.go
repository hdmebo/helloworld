package main

//some commets x

import (
	"fmt"
)

type stack []string

func Calc(input string) int {
	// commit 1
	// commit 2
	// commit 3
	// commit 4
	return 0
}

// change line - always return given string
func Calc2(input string) string {
	return input + "z"
}

func Calc3(input string) int {
	// commit 3
	return 3
}

func Calc4(input string) int {
	// commit 4
	return 4 + 2
}

func (s *stack) Push(value string) {
	fmt.Printf("BBB %p\n", s)
	temp := append(*s, value)
	*s = temp
	fmt.Printf("AAA %p %p\n", temp, s)

}
func (s *stack) Pop() string {
	fmt.Println("YYY ", *s, cap(*s), len(*s))
	if len(*s) == 0 {
		return ""
	}
	last := len(*s) - 1
	value := (*s)[last]
	*s = (*s)[:last]
	fmt.Println("XXX ", value, *s, cap(*s), len(*s))
	return value
}

// Calc1 is calc return 1
func Calc1(input string) int {
	return 1
}

func main() {
	stack := stack{}
	stack.Push("3")
	stack.Push("5")
	stack.Push("8")
	// operatorStack := []string{}
	fmt.Printf("%p %v\n", stack, stack)
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
