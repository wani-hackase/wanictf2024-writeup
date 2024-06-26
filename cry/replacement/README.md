---
title: replacement   
level: 2          
flag: FLAG{13epl4cem3nt}
writer: Gureisya      
---

# replacement

## 問題文
No one can read my diary!

## 解法
ハッシュ関数は内部状態が同じであれば同じ文字に対して同じ出力を返すので、1バイトの文字全ての入力と出力の組を記録した表を作ることで日記を復元出来ます。

## soluion
Hash functions produce the same output for the same input, so by recording input-output pairs for all 1-byte characters, one can reconstruct the original message.

## solver
```py
import hashlib

enc = None
with open('my_diary_11_8_Wednesday.txt') as f:
       enc = f.readline().strip()
       enc = eval(enc)

d = {}
for i in range(0, 256):
    x = hashlib.md5(str(i).encode()).hexdigest()
    num = int(x, 16)
    d[num] = i
    
dec = ''
for i in enc:
    dec += chr(d[i])
    
print(dec)
```