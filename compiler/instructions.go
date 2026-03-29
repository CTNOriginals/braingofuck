package compiler

type BlockType int

const (
	Head BlockType = iota

	ADV
	BAC
	INC
	DEC
	OUT
	INP
	BEG
	END

	STDIN
	Exit
)

const LoopPlaceholder = "#LOOP#"

var InstructionBlocks = map[BlockType][]string{
	Head: {
		"bits 64",
		"default rel",
		"global main",
		"extern GetStdHandle",
		"extern WriteFile",
		"extern ReadFile",
		"extern ExitProcess",
		"section .bss",
		"	tape 		resb 256",
		"	stdread		resq 1",
		"	stdwrite 	resq 1",
		"	written 	resb 4",
		"	input 		resb 256",
		"	input_len 	resq 1",
		"section .text",
		"main:",
		"	push r12",
		"	push r13",
		"	push r14",
		"	push r15",
		"	sub rsp, 56",
		"	mov rcx, -10",
		"	call GetStdHandle",
		"	mov [stdread], rax",
		"	mov rcx, -11",
		"	call GetStdHandle",
		"	mov [stdwrite], rax",
		"	lea r12, [tape]",
		"	xor r13, r13",
		"	lea r14, [input]",
		"	xor r15, r15",
	},

	ADV: {
		"inc r13",
	},
	BAC: {
		"dec r13",
	},
	INC: {
		"inc [r12 + r13]",
	},
	DEC: {
		"dec [r12 + r13]",
	},
	OUT: {
		"mov rcx, [stdwrite]",
		"lea rdx, [r12 + r13]",
		"mov r8,  1",
		"lea r9,  [written]",
		"mov qword [rsp + 32], 0",
		"call WriteFile",
	},
	INP: {
		"mov rcx, [r14 + r15]",
		"mov [r12 + r13], rcx",
		"inc r15",
	},
	BEG: {
		LoopPlaceholder + ":",
	},
	END: {
		"cmp [r12 + r13], 0",
		"jne " + LoopPlaceholder,
	},

	STDIN: {
		"mov rcx, [stdread]",
		"lea rdx, input",
		"mov r8, 64",
		"lea r9, input_len",
		"mov qword [rsp + 32], 0",
		"call ReadFile",
	},

	Exit: {
		"exit:",
		"	add rsp, 56",
		"	pop r15",
		"	pop r14",
		"	pop r13",
		"	pop r12",
		"	xor rcx, rcx",
		"	call ExitProcess",
	},
}
