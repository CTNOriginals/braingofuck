package interpreter

import "log"

type Stack []int

func CreateStack() *Stack {
	var stack = make(Stack, 0)
	return &stack
}

func (this *Stack) Push(val int) {
	*this = append(*this, val)
}

func (this *Stack) Pop() int {
	var stack = *this

	if len(stack) == 0 {
		log.Fatal("Can not pop empty stack")
	}

	var val = stack[len(stack)-1]

	*this = stack[0 : len(stack)-1]

	return val
}

func (this Stack) Peek() int {
	return this[len(this)-1]
}
