package compiler

type Cell byte

func (this *Cell) Inc() {
	*this += 1
}

func (this *Cell) Dec() {
	*this -= 1
}
