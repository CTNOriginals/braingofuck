package main

import (
	"fmt"
	"time"

	ctnfile "github.com/CTNOriginals/CTNGoUtils/v2/file"
	"github.com/CTNOriginals/braingofuck/compiler"
	"github.com/CTNOriginals/braingofuck/tokenizer"
)

func main() {
	var startTime = time.Now()
	fmt.Printf("\n\n---- BrainGoFuck START %s ----\n", startTime.Format(time.TimeOnly))
	defer func() {
		fmt.Printf("\n---- BrainGoFuck END %s (%f) ----\n", startTime.Format(time.TimeOnly), time.Since(startTime).Seconds())
	}()

	// var content = ctnfile.GetFileRunes("./brainfuck/proto.bf")
	var content = ctnfile.GetFileRunes("./brainfuck/hello.bf")
	var tokens = tokenizer.Tokenize(content)

	compiler.Compile(tokens, 256)
}
