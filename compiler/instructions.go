package compiler

import "github.com/CTNOriginals/braingofuck/tokenizer"

const (
	MOV = "mov"
	INC = "inc"
	DEC = "dec"
	SUB = "sub"
	LEA = "lea"
	XOR = "xor"

	PUSH = ""
	POP  = ""
	CALL = "call"

	TAPE string = "r12"
	OFFS string = "r13"
	INPT string = "r14"
	INPO string = "r15"

	VAL    string = "[" + TAPE + " + " + OFFS + "]"
	INPVAL string = "[" + INPT + " + " + INPO + "]"
)

type Instruction []string

var TokenInstructions = map[tokenizer.TokenType][]Instruction{
	tokenizer.ADV: {{
		"inc", TAPE,
	}},
	tokenizer.BAC: {{
		"dec", TAPE,
	}},
	tokenizer.INC: {{
		"inc", OFFS,
	}},
	tokenizer.DEC: {{
		"dec", OFFS,
	}},
	tokenizer.OUT: {
		{"mov", "rcx", "[stdread]"},
		{"lea", "rdx", "input"},
		{"mov", "r8", "4"},
		{"lea", "r9", "input_len"},
		{"mov", "qword [rsp + 32]", "0"},
		{"call", "ReadFile"},
	},
	tokenizer.INP: {
		{"mov", "rcx", INPVAL},
		{"mov", VAL, "rcx"},
		{"inc", INPO},
	},
	tokenizer.BEG: {{
		"",
	}},
	tokenizer.END: {{
		"",
	}},
}
