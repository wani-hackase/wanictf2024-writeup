---
title: "dance"
level: 3
flag: "FLAG{d4nc3_l0b0t_d4nc3!!}"
writer: "Gureisya"
---
# dance

## 問題文
step by step

## フラグ
`FLAG{d4nc3_l0b0t_d4nc3!!}`

## 解法
mycipher.pyを見るとquarter_roundという特徴的な関数があります。これを手掛かりに調べていくと、この暗号がChacha20であることが分かります。

Chacha20はkeyとnonceがあれば復号可能です。この問題ではどちらもトークンのハッシュ値の一部から生成されているためトークンを復元することが目標になります。トークンはユーザーネームとregisterした時の時刻、ランダムな0~10の値を用いて生成されているため、総当たりすることが可能です。

## solution
Based on the presence of the "quarter_round" function in the "mycipher.py" file, we can deduce that the cipher being used is ChaCha20.

ChaCha20 requires a key and a nonce to decrypt the message. In this case, both the key and nonce are derived from parts of the hash of the token. Therefore, the goal is to recover the token, which is generated using the user's username, the time of registration, and a random value between 0 and 10. This can be achieved through a brute-force approach.

## solver
```py
import hashlib
from mycipher import MyCipher

username = 'gureisya'
ciphertext = '3da5f9fa6998a991cb244a12fa72d311f3e6e9fbcac9984c0c'

def make_token(data1: str, data2: str):
    sha256 = hashlib.sha256()
    sha256.update(data1.encode())
    right = sha256.hexdigest()[:20]
    sha256.update(data2.encode())
    left = sha256.hexdigest()[:12]
    token = left + right
    return token

for sec in range(60):
    for minutes in range(60):
        data1 = f'user: {username}, {minutes}:{sec}'
        for i in range(10):
            data2 = f'{username}{i}'
            token = make_token(data1, data2)
            sha256 = hashlib.sha256()
            sha256.update(token[:32].encode())
            key = sha256.hexdigest()[:32]
            nonce = token[:12]
            cipher = MyCipher(key.encode(), nonce.encode())
            decrypted = cipher.encrypt(bytes.fromhex(ciphertext))
            if b'FLAG' in decrypted:
                print(decrypted)
                exit()
```
