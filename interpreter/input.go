package interpreter

type Input struct {
	Content []rune
	index   int
}

func CreateInput(content []rune) *Input {
	return &Input{
		Content: content,
	}
}

func (this *Input) Read() rune {
	var val rune = 0

	if this.index < len(this.Content) {
		val = this.Content[this.index]
	}

	this.index += 1

	return val
}
