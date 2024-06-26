---
title: "lambda"
level: 2
flag: "FLAG{l4_1a_14mbd4}"
writer: "hi120ki"
---

# lambda

## 問題文

Let's dance with lambda!

```
$ python lambda.py
Enter the flag:
```

## フラグ

`FLAG{l4_1a_14mbd4}`

## 解法

```python
def reverse_transformation(encoded):
    from pathlib import Path

    # ステップ1: 正解の文字列を分解し、各部分を基数36から整数に変換し、10を引く
    # Step1 : Split the correct string, convert each part from base 36 to integer, and subtract 10
    parts = encoded.split('_')
    decoded_chars = [chr(int(part, 36) + 10) for part in parts]

    # 結合して文字列に変換
    # Join and convert to string
    joined_string = ''.join(decoded_chars)

    # ステップ2: 123とXOR演算
    # Step2: XOR operation with 123
    xor_transformed = ''.join(chr(123 ^ ord(c)) for c in joined_string)

    # ステップ3: 各文字のASCIIコードに3を加算
    # Step3: Add 3 to the ASCII code of each character
    add_three = ''.join(chr(ord(c) + 3) for c in xor_transformed)

    # ステップ4: 各文字のASCIIコードから12を減算
    # Step4: Subtract 12 from the ASCII code of each character
    original_string = ''.join(chr(ord(c) - 12) for c in add_three)

    return original_string

encoded_string = '16_10_13_x_6t_4_1o_9_1j_7_9_1j_1o_3_6_c_1o_6r'
original_string = reverse_transformation(encoded_string)
print(original_string)
```
