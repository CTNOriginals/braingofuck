package main

import (
	"fmt"
	"time"

	"github.com/CTNOriginals/braingofuck/interpreter"
	"github.com/CTNOriginals/braingofuck/tokenizer"
)

var filePath string
var brainfuck = []rune{}
var input = []rune{}
var size = 256

func main() {
	var startTime = time.Now()
	fmt.Printf("\n\n---- BrainGoFuck START %s ----\n", startTime.Format(time.TimeOnly))
	defer func() {
		fmt.Printf("\n---- BrainGoFuck END %s (%f) ----\n", startTime.Format(time.TimeOnly), time.Since(startTime).Seconds())
	}()

	parseArgs()

	var tokens = tokenizer.Tokenize(brainfuck)

	interpreter.Inputpret(tokens, size, input)
}
