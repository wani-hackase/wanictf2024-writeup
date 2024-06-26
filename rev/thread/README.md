---
title: "Thread"
level: 4
flag: "FLAG{c4n_y0u_dr4w_4_1ine_be4ween_4he_thread3}"
writer: "hi120ki"
---

## 問題文

ワ...ワァ...!?

## 解法

ロックを使いつつ、スレッディングプログラミングを使っているバイナリを解析する問題です。
This is a challenge to analyze a binary that uses threading programming while using locks.

以下の3つの処理がIndexのModによって分類されて実行されます。
The following three processes are classified and executed by Index Mod.

```
inputs[index] = inputs[index] * 3;
inputs[index] = inputs[index] + 5;
inputs[index] = inputs[index] ^ 127;
```

これらの処理を把握し、それぞれの処理を逆変換することでフラグを取得することができます。
By understanding these processes and reversing each process, you can get the flag.

```python
a = [
    168,
    138,
    191,
    165,
    765,
    89,
    222,
    36,
    101,
    271,
    222,
    35,
    349,
    66,
    44,
    222,
    9,
    101,
    222,
    81,
    239,
    319,
    36,
    83,
    349,
    72,
    83,
    222,
    9,
    83,
    331,
    36,
    101,
    222,
    54,
    83,
    349,
    18,
    74,
    292,
    63,
    95,
    334,
    213,
    11,
]

for i in range(len(a)):
    if i % 3 == 0:
        a[i] = int(a[i] ^ 127)
        a[i] = int(a[i] - 5)
        a[i] = int(a[i] / 3)
    if i % 3 == 1:
        a[i] = int(a[i] / 3)
        a[i] = int(a[i] ^ 127)
        a[i] = int(a[i] - 5)
    if i % 3 == 2:
        a[i] = int(a[i] - 5)
        a[i] = int(a[i] / 3)
        a[i] = int(a[i] ^ 127)

flag = ""
for i in a:
    flag += chr(i)
print(flag)
```
