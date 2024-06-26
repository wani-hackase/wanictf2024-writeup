---
title: beginners_rsa   
level: 1          
flag: FLAG{S0_3a5y_1254!!}
writer: Gureisya      
---

# beginners_rsa

## 問題文

Do you know RSA?

## 解法
RSA暗号の公開鍵nが64ビットの5つの素数の積からなるため素因数分解することが出来ます。

素因数分解が完了したら秘密鍵dは次のように計算することが出来ます。

$$
\begin{aligned}
\phi = \sum (p_i-1)\space(p_i | \prod p_i = n)
\\
d \equiv e^{-1}\mod \phi
\end{aligned}
$$

あとは復号すればよく、以下の手順で平文mを求められます。

$$
\begin{aligned}
m \equiv c^{d}\mod n
\end{aligned}
$$

## solution
Since the public key n of the RSA cryptosystem is the product of five 64-bit prime numbers, it can be factored.

Once the prime factorization is completed, the secret key d can be calculated as follows:

$$
\begin{aligned}
\phi = \sum (p_i-1)\space(p_i | \prod p_i = n)
\\
d \equiv e^{-1}\mod \phi
\end{aligned}
$$

Then, to decrypt, the plaintext m can be recovered using:

$$
\begin{aligned}
m \equiv c^{d}\mod n
\end{aligned}
$$

## solver
```py
from Crypto.Util.number import *

n = 317903423385943473062528814030345176720578295695512495346444822768171649361480819163749494400347
e = 65537
enc = 127075137729897107295787718796341877071536678034322988535029776806418266591167534816788125330265

fs = factor(n)
phi = 1
for f in fs:
    phi *= f[0]-1
d = inverse(e, phi)
m = pow(enc, d, n)
print(long_to_bytes(int(m)))
```
