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
	lea r12, [tape]
	xor r13, r13
	lea r14, [input]
	xor r15, r15
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc r13
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc r13
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc [r12 + r13]
	inc r13
	inc [r12 + r13]
	inc [r12 + r13]
	dec r13
	dec r13
	dec r13
	loop0:
		mov rcx, [stdwrite]
		lea rdx, [r12 + r13]
		mov r8,  1
		lea r9,  [written]
		mov qword [rsp + 32], 0
		call WriteFile
		dec [r12 + r13]
		loop1:
			mov rcx, [stdwrite]
			lea rdx, [r12 + r13]
			mov r8,  1
			lea r9,  [written]
			mov qword [rsp + 32], 0
			call WriteFile
			dec [r12 + r13]
			loop2:
				mov rcx, [stdwrite]
				lea rdx, [r12 + r13]
				mov r8,  1
				lea r9,  [written]
				mov qword [rsp + 32], 0
				call WriteFile
				dec [r12 + r13]
				loop3:
					mov rcx, [stdwrite]
					lea rdx, [r12 + r13]
					mov r8,  1
					lea r9,  [written]
					mov qword [rsp + 32], 0
					call WriteFile
					dec [r12 + r13]
					cmp [r12 + r13], 0
					jne loop3
				inc r13
				cmp [r12 + r13], 0
				jne loop2
			inc r13
			cmp [r12 + r13], 0
			jne loop1
		inc r13
		cmp [r12 + r13], 0
		jne loop0
exit:
	add rsp, 56
	pop r15
	pop r14
	pop r13
	pop r12
	xor rcx, rcx
	call ExitProcess
