package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	ctnfile "github.com/CTNOriginals/CTNGoUtils/v2/file"
	"github.com/CTNOriginals/ctngoargs"
)

var arguments = ctngoargs.ArgDefinitionList{
	{Flags: []string{"--help"},
		Description: []string{
			"Prints a list of arguments and their descriptions.",
			"Follow it up with another argument (without the -- before it)",
			"to get a more detailed explanation about that argument.",
		},
		Inputs: ctngoargs.ArgInputs{
			{
				Name: "flag",
				Description: []string{
					"The name (without the -- before it) of a flag.",
					"Print a detailed explanation about a specific flag.",
				},
			},
		},
		Action: func(inputs []string) {
			fmt.Print(fakeargs.GetHelpString(inputs))
			os.Exit(0)
		},
	},
	{Flags: []string{"--full-help"},
		Description: []string{
			"Like --help, but include full detail for every flag.",
		},
		Action: func(_ []string) {
			fmt.Print(fakeargs.GetFullHelpString())
			os.Exit(0)
		},
	},
	{Flags: []string{"--brainfuck", "--bf"},
		Description: []string{
			"Pass in the brainfuck code to compile",
		},
		Inputs: ctngoargs.ArgInputs{
			{
				Name: "content",
				Description: []string{
					"The brainfuck code",
				},
			},
		},
		Action: func(inputs []string) {
			if len(inputs) == 0 {
				log.Fatalf("'--brainfuck' is missing the following input: 'content'")
			}

			brainfuck = []rune(inputs[0])
		},
	},
	{Flags: []string{"--file"},
		Description: []string{
			"Pass in a file to compile the contents of",
		},
		Inputs: ctngoargs.ArgInputs{
			{
				Name: "path",
				Description: []string{
					"The file path to compile.",
				},
			},
		},
		Action: func(inputs []string) {
			if len(inputs) == 0 {
				log.Fatalf("'--file' is missing the following input: 'path'")
			}

			var path = inputs[0]

			if !ctnfile.FileExists(path) {
				log.Fatalf("'--file' path does not exist: %s", path)
			}

			brainfuck = ctnfile.GetFileRunes(path)
		},
	},
	{Flags: []string{"--input"},
		Description: []string{
			"Set the input for the for the brainfuck program to receive.",
		},
		Inputs: ctngoargs.ArgInputs{
			{
				Name: "input",
				Description: []string{
					"The content of the input.",
				},
			},
		},
		Action: func(inputs []string) {
			if len(inputs) == 0 {
				log.Fatalf("'--input' is missing the following input: 'input' (yeah this is confusing...)")
			}

			input = []rune(inputs[0])
		},
	},
	{Flags: []string{"--size"},
		Description: []string{
			"Set the max amount of cells that ram should have.",
		},
		Inputs: ctngoargs.ArgInputs{
			{
				Name: "amount",
				Description: []string{
					"The number of memory cells.",
					"Default: 256",
				},
			},
		},
		Action: func(inputs []string) {
			if len(inputs) == 0 {
				log.Fatalf("'--size' is missing the following input: 'amount'")
			}

			var num, err = strconv.Atoi(inputs[0])

			if err != nil {
				log.Fatalf("'--size' unable to convert input string (%s) to int: %v", inputs[0], err)
			}

			size = num
		},
	},
}

var fakeargs ctngoargs.ArgDefinitionList

func parseArgs() {
	var args = os.Args
	fakeargs = arguments
	ctngoargs.ParseArgs(arguments, args)
}
