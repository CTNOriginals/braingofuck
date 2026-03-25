bits 64
default rel

global main
extern GetStdHandle
extern WriteFile
	; [in]                HANDLE       hFile,
	; [in]                LPCVOID      lpBuffer,
	; [in]                DWORD        nNumberOfBytesToWrite,
	; [out, optional]     LPDWORD      lpNumberOfBytesWritten,
	; [in, out, optional] LPOVERLAPPED lpOverlapped

extern ReadFile
	; [in]                HANDLE       hFile,
	; [out]               LPVOID       lpBuffer,
	; [in]                DWORD        nNumberOfBytesToRead,
	; [out, optional]     LPDWORD      lpNumberOfBytesRead,
	; [in, out, optional] LPOVERLAPPED lpOverlapped

extern ExitProcess

section .data
	msg			db "Please enter your name: "
	msg_len 	equ $ - msg
	greet 		db "Hello "
	greet_len 	equ $ - greet

section .bss
	stdread		resq 1
	stdwrite	resq 1
    written 	resq 1
	read 		resb 64
	read_len 	resq 1

section .text
main:
	; shadow space (32) + alignment + local var space
    sub rsp, 56

	mov rcx, -10
    call GetStdHandle
	mov [stdread], rax

	mov rcx, -11
    call GetStdHandle
	mov [stdwrite], rax

do_write:
    mov rcx, [stdwrite]
    lea rdx, [msg]
    mov r8,  msg_len
    lea r9,  [written]
    mov qword [rsp + 32], 0 ; 5th arg on stack
    call WriteFile

do_read:
	mov rcx, [stdread]
	lea rdx, read
	mov r8, 64
	lea r9, read_len
    mov qword [rsp + 32], 0 ; 5th arg on stack
	call ReadFile

do_greet:
    mov rcx, [stdwrite]
    lea rdx, [greet]
    mov r8,  greet_len
    lea r9,  [written]
    mov qword [rsp + 32], 0 ; 5th arg on stack
    call WriteFile

    mov rcx, [stdwrite]
    lea rdx, [read]
    mov r8,  [read_len]
    lea r9,  [written]
    mov qword [rsp + 32], 0 ; 5th arg on stack
    call WriteFile

exit: ; ExitProcess(0)
	jmp do_write ; inf loop
    xor rcx, rcx
    call ExitProcess
