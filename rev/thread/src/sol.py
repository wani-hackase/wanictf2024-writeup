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
