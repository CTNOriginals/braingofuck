package main

import (
	"fmt"
	"time"

	ctnfile "github.com/CTNOriginals/CTNGoUtils/v2/file"
	"github.com/CTNOriginals/braingofuck/tokenizer"
)

func main() {
	var startTime = time.Now()
	fmt.Printf("\n\n---- BrainGoFuck START %s ----\n", startTime.Format(time.TimeOnly))
	defer func() {
		fmt.Printf("---- BrainGoFuck END %s (%f) ----\n", startTime.Format(time.TimeOnly), time.Since(startTime).Seconds())
	}()

	var content = ctnfile.GetFileBytes("./brainfuck/proto.bf")
	var tokens = tokenizer.Tokenize(content)

	for i, token := range tokens {
		fmt.Printf("%d: %s %d:%d\n", i, tokenizer.TokenTypeValues[token.Typ], token.Line, token.Col)
	}
}
