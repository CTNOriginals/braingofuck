package compiler

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	ctnfile "github.com/CTNOriginals/CTNGoUtils/v2/file"
	ctnstring "github.com/CTNOriginals/CTNGoUtils/v2/string"
	"github.com/CTNOriginals/braingofuck/interpreter"
	"github.com/CTNOriginals/braingofuck/tokenizer"
)

func Compile(tokens tokenizer.TokenList, filePath string) {
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

	execute(destFile.Name())
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

func execute(destPath string) {
	var pathData = ctnfile.ParseFilePath(destPath)
	var binPath = fmt.Sprintf("%s/bin", pathData.Path)
	var binFile = fmt.Sprintf("%s/%s", binPath, pathData.Name)

	var exe = fmt.Sprintf("%s.exe", binFile)
	var obj = fmt.Sprintf("%s.obj", binFile)

	// fmt.Printf("destPath: \t%s\n", destPath)
	// fmt.Printf("binPath: \t%s\n", binPath)
	// fmt.Printf("binFile: \t%s\n", binFile)
	// fmt.Printf("exe: \t\t%s\n", exe)
	// fmt.Printf("obj: \t\t%s\n", obj)
	// fmt.Printf("std: \t\t%s\n", std)

	var nasm = []string{
		"nasm",
		"-f", "win64",
		"-X", "gcc",
		"--debug",
		"-g",
		"-o", obj,
		destPath,
	}
	var gcc = []string{"gcc", "-o", exe, obj}
	var run = []string{exe}

	runcmd(nasm)
	runcmd(gcc)
	var stdout = runcmd(run)

	fmt.Printf("-- COMPILER OUTPUT --\n%s\n%v\n", stdout, stdout)
}

func runcmd(args []string) []byte {
	var cmd = exec.Command(args[0], args[1:]...)
	var output, err = cmd.CombinedOutput()

	// fmt.Println(cmd.String())

	if err != nil {
		log.Fatalf("Error while executing command '%s' %v: %v", args[0], args[1:], err)
	}

	// fmt.Printf("%s output: %v\n", args[0], output)
	return output
}
