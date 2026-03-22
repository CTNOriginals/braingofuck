package interpreter

type Process struct {
	Ram    *Ram
	Stack  *Stack
	Stdin  *Input
	Stdout []Cell
}

func CreateProcess(size int, input []rune) *Process {
	return &Process{
		Ram:    CreateRam(size),
		Stack:  CreateStack(),
		Stdin:  CreateInput(input),
		Stdout: make([]Cell, 0),
	}
}
