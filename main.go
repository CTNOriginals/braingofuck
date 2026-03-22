package main

import (
	"fmt"
	"time"

	"github.com/CTNOriginals/braingofuck/interpreter"
	"github.com/CTNOriginals/braingofuck/tokenizer"
)

var brainfuck = []rune{}
var input = []rune{}

func main() {
	var startTime = time.Now()
	fmt.Printf("\n\n---- BrainGoFuck START %s ----\n", startTime.Format(time.TimeOnly))
	defer func() {
		fmt.Printf("\n---- BrainGoFuck END %s (%f) ----\n", startTime.Format(time.TimeOnly), time.Since(startTime).Seconds())
	}()

	parseArgs()

	var tokens = tokenizer.Tokenize(brainfuck)

	interpreter.Inputpret(tokens, 256, input)
}
