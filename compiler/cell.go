package compiler

type Cell rune

func (this *Cell) Set(val rune) {
	*this = Cell(val)
}

func (this *Cell) Inc() {
	*this += 1
}

func (this *Cell) Dec() {
	*this -= 1
}
