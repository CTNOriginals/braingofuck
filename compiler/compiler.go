package compiler

import (
	"fmt"

	"github.com/CTNOriginals/braingofuck/tokenizer"
)

var ram *Ram
var stack *Stack
var stdout []Cell

func Compile(tokens tokenizer.TokenList, size int) {
	ram = CreateRam(size)
	stack = CreateStack()
	for cursor := 0; cursor < len(tokens); cursor++ {
		var token = tokens[cursor]

		if token.Col == 1 {
			fmt.Printf("-- %d:%d (%d) --\n", token.Line, token.Col, cursor)
			for adr := range 3 {
				fmt.Printf("%d: %d\n", adr, *ram.Mem[adr])
			}
			fmt.Printf("Stack: %v\n", *stack)
		}

		switch token.Typ {
		case tokenizer.ADV:
			ADV()
			BAC()
			INC()
		case tokenizer.DEC:
			DEC()
		case tokenizer.INP:
			INP()
		case tokenizer.OUT:
			OUT()
		case tokenizer.BEG:
			BEG(cursor)
		case tokenizer.END:
			var before = cursor
			cursor = END(cursor)
			fmt.Printf("%d > %d ", before, cursor)
		}
	}

	println("\n-- OUT --")
	fmt.Printf("%v\n", stdout)
}

func ADV() {
func BAC() {
	ram.Backup()
}
func INC() {
func DEC() {
	ram.Dec()
}
func INP() {
	// println("TODO: Compiler INP")
}
func OUT() {
	stdout = append(stdout, *ram.Value())
}
func BEG(index int) {
	stack.Push(index)
	fmt.Printf("Push: %d (%v)\n", index, *stack)
}
func END(index int) int {
	var val = ram.Value()

	fmt.Printf(" val:%d ", *val)

	if *val != 0 {
		return stack.Peek()
	}

	fmt.Printf("\ni:%d pop:%d (%v)\n", index, stack.Pop(), *stack)

	return index
}
