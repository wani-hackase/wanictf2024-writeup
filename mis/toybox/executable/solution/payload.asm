    global _start

    section .text
_start:
    ; read(7, buf, 256)
    mov rax, 0      ; syscall #0 (read)
    mov rdi, 7      ; from fd 7
    mov rsi, buf    ; to buf
    mov rdx, 256    ; number of bytes
    syscall

    ; write(1, buf, 256)
    mov rax, 1      ; syscall #1 (write)
    mov rdi, 1      ; to fd 1 (stdout)
    mov rsi, buf    ; from buf
    mov rdx, 256    ; number of bytes
    syscall

    ; exit(0)
    mov rax, 60     ; syscall #60 (exit)
    mov rdi, 0      ; code 0
    syscall

    section .bss
buf:
    resb 256
