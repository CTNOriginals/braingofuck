package compiler

type Ram struct {
	Mem     []*Cell
	Address int
	size    int
}

func CreateRam(size int) *Ram {
	var mem = make([]*Cell, size)

	for i := range size {
		mem[i] = new(Cell)
	}

	return &Ram{
		Mem:     mem,
		Address: 0,
		size:    size,
	}
}

func (this Ram) Value() *Cell {
	return this.Mem[this.Address]
}

func (this *Ram) Advance() {
	this.Address += 1

	if this.Address == this.size {
		this.Address = 0
	}
}

func (this *Ram) Backup() {
	this.Address -= 1

	if this.Address < 0 {
		this.Address += this.size
	}
}

func (this *Ram) Inc() {
	(*this).Mem[this.Address].Inc()
}

func (this *Ram) Dec() {
	(*this).Mem[this.Address].Dec()
}
