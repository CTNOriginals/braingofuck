package compiler

type Assembly []string

func (this *Assembly) Insert(lines []string) {
	*this = append(*this, lines...)
}

func (this Assembly) AsBytes() []byte {
	var content = []byte{}

	for _, line := range this {
		content = append(content, []byte(line)...)
		content = append(content, '\n')
	}

	return content
}
