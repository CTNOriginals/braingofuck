bits 64
default rel

global main
extern GetStdHandle
extern WriteFile
extern ReadFile
extern ExitProcess

section .bss
	tape 		resb 256
	stdread		resq 1
	stdwrite 	resq 1
	written 	resb 4
	input 		resb 256
	input_len 	resq 1

section .text
main:
	; Init
	push r12
	push r13
	push r14
	push r15
	sub rsp, 56

	mov rcx, -10
	call GetStdHandle
	mov [stdread], rax

	mov rcx, -11
	call GetStdHandle
	mov [stdwrite], rax

	lea r12, [tape] ; tape start
	xor r13, r13 	; offset

	lea r14, [input] ; input start
	xor r15, r15 	; input offset

	; Program
	; ,[.,]

	;,
	mov rcx, [stdread]
	lea rdx, input
	mov r8, 64
	lea r9, input_len
	mov qword [rsp + 32], 0
	call ReadFile

	mov rcx, [r14 + r15]
	mov [r12 + r13], rcx
	inc r15

	;[.,]
	loop1:
		mov rcx, [stdwrite]
		lea rdx, [r12 + r13]
		mov r8,  1
		lea r9,  [written]
		mov qword [rsp + 32], 0
		call WriteFile

		mov rcx, [r14 + r15]
		mov [r12 + r13], rcx
		inc r15

		cmp [r12 + r13], 0
	jne loop1

exit:
	add rsp, 56
    pop r15
    pop r14
    pop r13
    pop r12
	xor rcx, rcx
	call ExitProcess
