package compiler

import "log"

type Stack []int

func (this *Stack) Push(val int) {
	*this = append(*this, val)
}

func (this *Stack) Pop() int {
	if len((*this)) == 0 {
		log.Fatal("Can not pop empty stack")
	}

	var val = (*this)[len(*this)-1]

	(*this) = (*this)[0 : len(*this)-1]

	return val
}

func (this Stack) Peek() int {
	return this[len(this)-1]
}
