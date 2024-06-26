---
title: Easy calc
level: 2
flag: FLAG{Do_the_math396691ba7d7270a}
writer: Laika
---

# Easy calc

## 問題文

😆

## 解法
フラグが $s$ の値を鍵として暗号化されています。
与えられた $A := f(s, p), p$ の値から $s$ を復元できないか考えてみます。
まず、 $A = f(s, p)$ の中身を式に起こしてみます。

$$
\begin{aligned}
A = f(s, p) &= s(\dots( s(s(s(p-0)+(p-1))+(p-2)) + \dots + 2) + 1 \mod p \\\\
  &= \sum_{i=0}^{p-1} s^{p-i}(p-i) \\\\
  &= \sum_{i=1}^{p} s^{i}i 
\end{aligned}
$$

さらに書き下し、 $sA$ との差を考えると

$$
\begin{aligned}
    A &= 1 \cdot s^1 + & 2 \cdot s^2 & + \cdots + & (p-1) \cdot s^{p-1} & + & p \cdot s^{p} & \\\\
-) sA &=               & 1 \cdot s^2 & + \cdots + & (p-2) \cdot s^{p-1} & + & (p-1) \cdot s^{p} & + p \cdot s^{p+1} \\\\
(1-s)A &= 1 \cdot s^1 + & 1 \cdot s^2 & + \cdots + & 1 \cdot s^{p-1} & + & 1 \cdot s^{p} & - p \cdot s^{p+1} \\\\
       &= \frac{s(1-s^p)}{1-s} \\\\
       &= s \ \ (\because s^p = s \mod p) \\\\
\therefore \ \ s &= \frac{A}{1+A} \mod p
\end{aligned}
$$

$s$ を $A, p$ の式で表せたので、これを鍵としてAES-CBCの暗号文を復号すればよいです。

### 補足
Cryptoでは往々にして数式をこねくり回すことがあるので、初心者の方にはこの工程に慣れてほしいという思いで作問しました。


## Solution 


The flag is encrypted using the value of $s$ as a key.
Let's consider if we can recover $s$ from the given values $A := f(s, p)$ and $p$.
First, let's express the contents of $A$ in equation form.

$$
\begin{aligned}
A = f(s, p) &= s(\dots( s(s(s(p-0)+(p-1))+(p-2)) + \dots + 2) + 1 \mod p \\\\
  &= \sum_{i=0}^{p-1} s^{p-i}(p-i) \\\\
  &= \sum_{i=1}^{p} s^{i}i 
\end{aligned}
$$

By further expanding and considering the difference with $sA$, we have

$$
\begin{aligned}
    A &= 1 \cdot s^1 + & 2 \cdot s^2 & + \cdots + & (p-1) \cdot s^{p-1} & + & p \cdot s^{p} & \\\\
-) sA &=               & 1 \cdot s^2 & + \cdots + & (p-2) \cdot s^{p-1} & + & (p-1) \cdot s^{p} & + p \cdot s^{p+1} \\\\
(1-s)A &= 1 \cdot s^1 + & 1 \cdot s^2 & + \cdots + & 1 \cdot s^{p-1} & + & 1 \cdot s^{p} & - p \cdot s^{p+1} \\\\
       &= \frac{s(1-s^p)}{1-s} \\\\
       &= s \ \ (\because s^p = s \mod p) \\\\
\therefore \ \ s &= \frac{A}{1+A} \mod p
\end{aligned}
$$

Since $s$ is expressed as a function of $A$ and $p$, we can decrypt the ciphertext encrypted with AES-CBC using this as the key.

In cryptography challenge, manipulating mathematical expressions is often required. 
This challenge was designed to help beginners get used to this kind of process.


### Solver
```python
import ast
from hashlib import md5

from Crypto.Cipher import AES
from Crypto.Util.number import *


def decrypt(ciphertext: bytes, key: int, iv: bytes) -> bytes:
    key = long_to_bytes(key)
    key = md5(key).digest()
    cipher = AES.new(key, AES.MODE_CBC, iv=iv)
    return cipher.decrypt(ciphertext)


with open("output.txt") as f:
    p = ast.literal_eval(f.readline().split("=")[1].strip())
    A = ast.literal_eval(f.readline().split("=")[1].strip())
    ciphertext = ast.literal_eval(f.readline().split("=")[1].strip())
    ciphertext = bytes.fromhex(ciphertext)


s = A * pow(1 + A, -1, p) % p
iv, ciphertext = ciphertext[:16], ciphertext[16:]
flag = decrypt(ciphertext, s, iv)
print(flag.decode())
```


