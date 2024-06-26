---
title: beginners_aes    
level: 1          
flag: FLAG{7h3_f1r57_5t3p_t0_Crypt0!!}
writer: Gureisya      
---

# beginners_aes

## 問題文

AES is one of the most important encryption methods in our daily lives.

## 解法
AESで暗号化されたデータは鍵と初期化ベクトルが分かっていれば復号することが出来ます。どちらもほとんど分かっているので残り1バイトを総当たりすることで解けます。

## solution
AES encryption can be decrypted if the key and initialization vector (IV) are known. In this case, since the key and most of the IV are known, the remaining 1 byte of the IV can be brute-forced to decrypt the data.

## solver
```py
from Crypto.Cipher import AES
from Crypto.Util.Padding import unpad
import hashlib

enc = b'\x16\x97,\xa7\xfb_\xf3\x15.\x87jKRaF&"\xb6\xc4x\xf4.K\xd77j\xe5MLI_y\xd96\xf1$\xc5\xa3\x03\x990Q^\xc0\x17M2\x18'
flag_hash = "6a96111d69e015a07e96dcd141d31e7fc81c4420dbbef75aef5201809093210e"
key = b'the_enc_key_is_'
iv = b'my_great_iv_is_'

for b1 in range(256):
    for b2 in range(256):
        key_ = key + bytes([b1])
        iv_ = iv + bytes([b2])
        cipher = AES.new(key_, AES.MODE_CBC, iv_)
        dec = cipher.decrypt(enc)
        try:
            m = unpad(dec, 16)
            m_hash = hashlib.sha256(m).hexdigest()
            if m_hash == flag_hash:
                print(f'decrypted = {m}')
        except:
            pass
```
