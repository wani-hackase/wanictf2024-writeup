---
title: speedy   
level: 4          
flag: FLAG{x013_ro74te_5hif7!!}
writer: Gureisya      
---

# speedy

## 問題文
I made a super speedy keystream cipher!!

## 解法
cipher.pyを見ると、二つの整数（内部状態）をxorやrotationを用いて更新しつつ鍵ストリームを生成してそれらとのxorを行うような暗号であることが分かります。この疑似乱数生成手法はXoroshiro128+と呼ばれます。

内部状態(X, Y)の更新は以下のように表現されます。

$$
\begin{aligned}
X = (rotl(X, 24) \oplus X \oplus Y \oplus ((X\oplus Y) << 16)) \mod 2^{64}-1
\\
Y = rotl((X \oplus Y), 37) \mod 2^{64}-1
\end{aligned}
$$

この操作はshiftとxorのみを用いているため行列として表すことが出来、かつ可逆な変換となっています。
よって、逆行列を計算することで内部状態更新の逆操作が求まります。以下が逆操作です。

$$
\begin{aligned}
X = rotl(X,40) \oplus rotl(Y,3) \oplus rotl((rotl(Y,27)<<16),40) \mod 2^{64}-1
\\
Y = rotl(X,40) \oplus rotl(Y,3) \oplus rotl((rotl(Y,27)<<16),40) \oplus rotl(Y,27) \mod 2^{64}-1
\end{aligned}
$$

あとはシードを復元する作業ですが、Xは暗号文に埋め込まれているためそのまま取り出したらよく、Yはパディングが大きいことを利用すれば256通りに絞ることが出来ます。

以上より、シードが分かったので鍵ストリームを生成すれば復号できます。

## solution
When looking at cipher.py, it becomes clear that this is a cipher that generates a key stream by updating two integers (internal states) using XOR and rotation operations, and then performing an XOR with those states. This pseudo-random number generation technique is called Xoroshiro128+.

The update of the internal state (X, Y) is expressed as follows:

$$
\begin{aligned}
X = (rotl(X, 24) \oplus X \oplus Y \oplus ((X\oplus Y) << 16)) \mod 2^{64}-1
\\
Y = rotl((X \oplus Y), 37) \mod 2^{64}-1
\end{aligned}
$$

This operation uses only shifts and XOR, and can be represented as a matrix, which is also an invertible transformation. Therefore, the inverse operation of the internal state update can be obtained by calculating the inverse matrix. The inverse operation is as follows:

$$
\begin{aligned}
X = rotl(X,40) \oplus rotl(Y,3) \oplus rotl((rotl(Y,27)<<16),40) \mod 2^{64}-1
\\
Y = rotl(X,40) \oplus rotl(Y,3) \oplus rotl((rotl(Y,27)<<16),40) \oplus rotl(Y,27) \mod 2^{64}-1
\end{aligned}
$$

The remaining task is to restore the seed. Since X is embedded in the ciphertext, it can be directly extracted, and Y can be narrowed down to 256 possibilities by utilizing the large padding.

With the seed known, the key stream can be generated, and decryption can be performed.

## solver
```py
from Crypto.Util.number import *
from Crypto.Util.Padding import *
import os

def rotl(x, y):
    x &= 0xFFFFFFFFFFFFFFFF
    return ((x << y) | (x >> (64 - y))) & 0xFFFFFFFFFFFFFFFF

class MyCipher:
    def __init__(self, s0, s1):
        self.X = s0
        self.Y = s1
        self.mod = 0xFFFFFFFFFFFFFFFF
        self.BLOCK_SIZE = 8
    
    def get_key_stream(self):
        s0 = self.X
        s1 = self.Y
        sum = (s0 + s1) & self.mod
        s1 ^= s0
        key = []
        for _ in range(8):
            key.append(sum & 0xFF)
            sum >>= 8
        
        self.X = (rotl(s0, 24) ^ s1 ^ (s1 << 16)) & self.mod
        self.Y = rotl(s1, 37) & self.mod
        return key
    
    def set_seed(self, s0, s1):
        self.X = s0
        self.Y = s1
    
    def encrypt(self, pt: bytes):
        ct = b''
        for i in range(0, len(pt), self.BLOCK_SIZE):
            ct += long_to_bytes(self.X)
            key = self.get_key_stream()
            block = pt[i:i+self.BLOCK_SIZE]
            ct += bytes([block[j] ^ key[j] for j in range(len(block))])
        return ct
    
    def decrypt(self, ct: bytes, s0=None, s1=None):
        pt = b''
        if s0 is not None and s1 is not None:
            self.set_seed(s0, s1)
        for i in range(0, len(ct), 2*self.BLOCK_SIZE):
            key = self.get_key_stream()
            block = ct[i+self.BLOCK_SIZE:i+2*self.BLOCK_SIZE]
            pt += bytes([block[j] ^ key[j] for j in range(len(block))])
        return unpad(pt, self.BLOCK_SIZE)
    
def rotl(x, y):
    x &= 0xFFFFFFFFFFFFFFFF
    return ((x << y) | (x >> (64 - y))) & 0xFFFFFFFFFFFFFFFF

def back(x, y, times):
    s0 = x
    s1 = y
    for _ in range(times):
        s0 = (rotl(x, 40) ^ rotl(y, 3) ^ rotl((rotl(y, 27)<<16), 40)) & 0xFFFFFFFFFFFFFFFF
        s1 = (rotl(x, 40) ^ rotl(y, 3) ^ rotl((rotl(y, 27)<<16), 40) ^ rotl(y, 27)) & 0xFFFFFFFFFFFFFFFF
        x = s0
        y = s1
    return s0, s1

ct = b'"G:F\xfe\x8f\xb0<O\xc0\x91\xc8\xa6\x96\xc5\xf7N\xc7n\xaf8\x1c,\xcb\xebY<z\xd7\xd8\xc0-\x08\x8d\xe9\x9e\xd8\xa51\xa8\xfbp\x8f\xd4\x13\xf5m\x8f\x02\xa3\xa9\x9e\xb7\xbb\xaf\xbd\xb9\xdf&Y3\xf3\x80\xb8'
flag = b''
for b in range(0x00, 0x100):
    sum = 0
    pred = long_to_bytes(b)+b'\x07'*7
    for i in range(8):
        sum += (ct[-8+i] ^ pred[i]) << 8*i

    x = bytes_to_long(ct[-16:-8])
    y = (sum-x) & 0xFFFFFFFFFFFFFFFF
    seed0, seed1 = back(x, y, 3)
    cipher = MyCipher(0xdeadbeef, 0x12345678)
    pt = cipher.decrypt(ct, seed0, seed1)
    if pt.startswith(b'FLAG{'):
        flag = pt
        break

print(flag)
```
