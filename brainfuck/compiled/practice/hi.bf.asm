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
	read 		resb 1
	read_len 	resq 1

section .text
main:
	; Init
	push r12
	push r13
	sub rsp, 56

	mov rcx, -10
	call GetStdHandle
	mov [stdread], rax

	mov rcx, -11
	call GetStdHandle
	mov [stdwrite], rax

	lea r12, [tape]
	xor r13, r13

	; Program
	;>>
	inc r13
	inc r13

	;++++++++++
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]

	;[<<++++++++++>>-]
	loop1:
		dec r13
		dec r13
		inc [r12 + r13]
		inc [r12 + r13]
		inc [r12 + r13]
		inc [r12 + r13]
		inc [r12 + r13]
		inc [r12 + r13]
		inc [r12 + r13]
		inc [r12 + r13]
		inc [r12 + r13]
		inc [r12 + r13]
		inc r13
		inc r13

		dec [r12 + r13]
		cmp [r12 + r13], 0
	jne loop1

	;++++++++++
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]

	;[<++++++++++>-]
	loop2:
		dec r13
		inc [r12 + r13]
		inc [r12 + r13]
		inc [r12 + r13]
		inc [r12 + r13]
		inc [r12 + r13]
		inc [r12 + r13]
		inc [r12 + r13]
		inc [r12 + r13]
		inc [r12 + r13]
		inc [r12 + r13]
		inc r13

		dec [r12 + r13]
		cmp [r12 + r13], 0
	jne loop2

	;<+++++
	dec r13

	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]

	;<++++
	dec r13

	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]

	;.>.
	mov rcx, [stdwrite]
    lea rdx, [r12 + r13]
    mov r8,  1
    lea r9,  [written]
    mov qword [rsp + 32], 0
    call WriteFile

	inc r13

	mov rcx, [stdwrite]
    lea rdx, [r12 + r13]
    mov r8,  1
    lea r9,  [written]
    mov qword [rsp + 32], 0
    call WriteFile

exit:
	add rsp, 56
    pop r13
    pop r12
	xor rcx, rcx
	call ExitProcess
