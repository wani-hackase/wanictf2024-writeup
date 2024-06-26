from cipher import MyCipher
from Crypto.Util.number import *
from Crypto.Util.Padding import *
import os

s0 = bytes_to_long(os.urandom(8))
s1 = bytes_to_long(os.urandom(8))

cipher = MyCipher(s0, s1)
secret = b'FLAG{'+b'*'*19+b'}'
pt = pad(secret, 8)
ct = cipher.encrypt(pt)
print(f'ct = {ct}')