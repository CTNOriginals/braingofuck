bits 64
default rel

global main
extern GetStdHandle
extern WriteFile
extern ReadFile
extern ExitProcess

section .bss
    written 	resq 1
	read 		resb 64
	read_len 	resq 1

section .text
main:
	; shadow space (32) + alignment + local var space
    sub rsp, 56

do_read:
    mov rcx, -10
    call GetStdHandle

	mov rcx, rax
	lea rdx, read
	mov r8, 64
	lea r9, read_len
	mov qword [rsp + 32], 0
	call ReadFile

do_write:
    mov rcx, -11
    call GetStdHandle

    ; WriteFile(handle, msg, len, &written, NULL)
    mov rcx, rax
    lea rdx, [read]
    mov r8,  [read_len]
    lea r9,  [written]
    mov qword [rsp + 32], 0         ; 5th arg on stack
    call WriteFile

exit: ; ExitProcess(0)
	jmp do_read ; inf loop
    xor rcx, rcx
    call ExitProcess
