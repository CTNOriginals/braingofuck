package interpreter

import (
	"fmt"

	"github.com/CTNOriginals/braingofuck/tokenizer"
)

func Inputpret(tokens tokenizer.TokenList, size int, input []rune) {
	var process = CreateProcess(size, input)

	for cursor := 0; cursor < len(tokens); cursor++ {
		var token = tokens[cursor]

		switch token.Typ {
		case tokenizer.ADV:
			process.Ram.Advance()
		case tokenizer.BAC:
			process.Ram.Backup()
		case tokenizer.INC:
			process.Ram.Inc()
		case tokenizer.DEC:
			process.Ram.Dec()
		case tokenizer.INP:
			process.Ram.Set(process.Stdin.Read())
		case tokenizer.OUT:
			process.Stdout = append(process.Stdout, *process.Ram.Get())
		case tokenizer.BEG:
			process.Stack.Push(cursor)
		case tokenizer.END:
			if *process.Ram.Get() == 0 {
				process.Stack.Pop()
				break
			}

			cursor = process.Stack.Peek()
		}
	}

	println("\n-- OUT --")
	println(string(process.Stdout))
	fmt.Printf("%v\n", process.Stdout)
}
