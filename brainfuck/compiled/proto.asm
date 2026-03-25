bits 64
default rel

global main
extern GetStdHandle
extern WriteFile
extern ExitProcess

section .data
    msg db "Hello, World!", 10
    msg_len equ $ - msg

section .bss
    written resq 1
	read 	resb 64

section .text
main:
    sub rsp, 56                     ; shadow space (32) + alignment + local var space

    ; GetStdHandle(-11) -> stdout handle
    mov rcx, -11
    call GetStdHandle

    ; WriteFile(handle, msg, len, &written, NULL)
    mov rcx, rax
    lea rdx, [msg]
    mov r8,  msg_len
    lea r9,  [written]
    mov qword [rsp + 32], 0         ; 5th arg on stack
    call WriteFile

    ; ExitProcess(0)
    xor rcx, rcx
    call ExitProcess
