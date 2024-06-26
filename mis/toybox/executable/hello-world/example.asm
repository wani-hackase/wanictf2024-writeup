    global _start

    section .text
_start:
    ; write(1, msg, 13)
    mov rax, 1      ; syscall #1 (write)
    mov rdi, 1      ; to fd 1 (stdout)
    mov rsi, msg    ; from msg
    mov rdx, 13     ; number of bytes
    syscall

    ; exit(0)
    mov rax, 60     ; syscall #60 (exit)
    mov rdi, 0      ; code 0
    syscall

    section .data
msg:
    db "Hello, World", 10
