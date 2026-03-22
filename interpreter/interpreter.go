package interpreter

import (
	"fmt"

	"github.com/CTNOriginals/braingofuck/tokenizer"
)

var ram *Ram
var stack *Stack
var stdin *Input
var stdout []Cell

func Inputpret(tokens tokenizer.TokenList, size int, input []rune) {
	ram = CreateRam(size)
	stack = CreateStack()
	stdout = make([]Cell, 0)
	stdin = CreateInput(input)

	for cursor := 0; cursor < len(tokens); cursor++ {
		var token = tokens[cursor]

		switch token.Typ {
		case tokenizer.ADV:
			ADV()
		case tokenizer.BAC:
			BAC()
		case tokenizer.INC:
			INC()
		case tokenizer.DEC:
			DEC()
		case tokenizer.INP:
			INP(input)
		case tokenizer.OUT:
			OUT()
		case tokenizer.BEG:
			BEG(cursor)
		case tokenizer.END:
			cursor = END(cursor)
		}
	}

	println("\n-- OUT --")
	println(string(stdout))
	fmt.Printf("%v\n", stdout)
}

func ADV() {
	ram.Advance()
}
func BAC() {
	ram.Backup()
}
func INC() {
	ram.Inc()
}
func DEC() {
	ram.Dec()
}
func INP(input []rune) {
	ram.Set(stdin.Read())
}
func OUT() {
	stdout = append(stdout, *ram.Get())
}
func BEG(index int) {
	stack.Push(index)
}
func END(index int) int {
	var val = ram.Get()

	if *val != 0 {
		return stack.Peek()
	}

	stack.Pop()

	return index
}
