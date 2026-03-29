package compiler

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	ctnfile "github.com/CTNOriginals/CTNGoUtils/v2/file"
	ctnstring "github.com/CTNOriginals/CTNGoUtils/v2/string"
	"github.com/CTNOriginals/braingofuck/interpreter"
	"github.com/CTNOriginals/braingofuck/tokenizer"
)

func Compile(tokens tokenizer.TokenList, filePath string) string {
	var destFile = createDestFile(filePath)
	var asm = Assembly{}
	var stdin = false
	var loopCount = 0
	var loopStack = interpreter.CreateStack()
	var indent = 1

	asm.Insert(InstructionBlocks[Head])

	for _, token := range tokens {
		var typ = token.Typ

		if typ == tokenizer.INP && !stdin {
			asm.Insert(InstructionBlocks[STDIN])
			stdin = true
		}

		var block, exists = InstructionBlocks[BlockType(typ)]

		if !exists {
			log.Fatalf("Tokentype '%s' is not defined as an instruction block", tokenizer.TokenTypeValues[typ])
		}

		var lines = make([]string, len(block))

		for i, line := range block {
			lines[i] = ctnstring.Indent(line, indent, "\t")
		}

		if typ == tokenizer.BEG {
			var label = fmt.Sprintf("loop%d", loopCount)
			lines = setLoopLabel(lines, label)
			loopStack.Push(loopCount)
			loopCount += 1
			indent += 1
		}

		if typ == tokenizer.END {
			var label = fmt.Sprintf("loop%d", loopStack.Pop())
			lines = setLoopLabel(lines, label)
			indent -= 1
		}

		asm.Insert(lines)
	}

	asm.Insert(InstructionBlocks[Exit])

	var _, err = destFile.Write(asm.AsBytes())

	if err != nil {
		log.Fatalf("Error while writing to the destination file: %v", err)
	}

	return destFile.Name()
}

func createDestFile(filePath string) *os.File {
	var absolutePath, _ = filepath.Abs(filePath)
	var pathData = ctnfile.ParseFilePath(absolutePath)
	var destFilePath = fmt.Sprintf("%s/compiled/%s.asm", pathData.Path, pathData.File)

	var destFile, err = os.Create(destFilePath)

	if err != nil {
		log.Fatalf("Unable to create destination file: %v", err)
	}

	return destFile
}

func setLoopLabel(block []string, label string) []string {
	var lines = make([]string, len(block))

	for i, line := range block {
		lines[i] = strings.ReplaceAll(line, LoopPlaceholder, label)
	}

	return lines
}
