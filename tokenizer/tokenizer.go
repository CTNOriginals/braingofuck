package tokenizer

type TokenType int
type Token struct {
	Typ  TokenType
	Line int
	Col  int
}
type TokenList []Token

const (
	NUL TokenType = iota

	ADV // >
	BAC // <
	INC // +
	DEC // -
	OUT // .
	INP // ,
	BEG // [
	END // ]
)

var TokenTypeValues = []string{"NUL", "ADV", "BAC", "INC", "DEC", "OUT", "INP", "BEG", "END"}

var charToken = map[rune]TokenType{
	'>': ADV,
	'<': BAC,
	'+': INC,
	'-': DEC,
	'.': OUT,
	',': INP,
	'[': BEG,
	']': END,
}

func Tokenize(content []rune) TokenList {
	var list = make(TokenList, 0)

	var line = 1
	var col = 1

	for _, char := range content {
		var typ, exist = charToken[char]

		if exist {
			var token = Token{
				Typ:  typ,
				Line: line,
				Col:  col,
			}

			list = append(list, token)
		}

		col += 1

		if char == '\n' {
			line += 1
			col = 1
		}
	}

	return list
}
