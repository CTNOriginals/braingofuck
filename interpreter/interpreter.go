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
			ram.Advance()
		case tokenizer.BAC:
			ram.Backup()
		case tokenizer.INC:
			ram.Inc()
		case tokenizer.DEC:
			ram.Dec()
		case tokenizer.INP:
			ram.Set(stdin.Read())
		case tokenizer.OUT:
			stdout = append(stdout, *ram.Get())
		case tokenizer.BEG:
			stack.Push(cursor)
		case tokenizer.END:
			if *ram.Get() == 0 {
				stack.Pop()
				break
			}

			cursor = stack.Peek()
		}
	}

	println("\n-- OUT --")
	println(string(stdout))
	fmt.Printf("%v\n", stdout)
}
