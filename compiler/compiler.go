package compiler

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	ctnfile "github.com/CTNOriginals/CTNGoUtils/v2/file"
	"github.com/CTNOriginals/braingofuck/tokenizer"
)

func Compile(tokens tokenizer.TokenList, filePath string) {
	var destFile = createDestFile(filePath)

	_ = destFile
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
