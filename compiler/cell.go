package compiler

type Cell byte

func (this *Cell) Inc() *Cell {
	*this += 1
	return this
}

func (this *Cell) Dec() *Cell {
	*this += 1
	return this
}
