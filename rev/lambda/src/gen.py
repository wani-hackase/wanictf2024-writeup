import sys

sys.setrecursionlimit(10000000)

# obfuscate

flag = "FLAG{l4_1a_14mbd4}"

t = flag  # input("Enter the flag: ")
p = lambda x: "".join(chr(ord(c) + 12) for c in x)
s = "".join(chr(ord(c) - 3) for c in p(t))
q = lambda x: "".join(chr(123 ^ ord(c)) for c in x)
u = "".join(q(c) for c in s)


def int_to_base36(num):
    num = num - 10

    if num < 0:
        raise ValueError("Negative numbers are not supported.")
    if num == 0:
        return "0"

    digits = "0123456789abcdefghijklmnopqrstuvwxyz"
    result = []

    while num:
        num, rem = divmod(num, 36)
        result.append(digits[rem])

    return "".join(reversed(result))


assert len(flag) == len(u)
res = []
for i in u:
    res.append(int_to_base36(ord(i)))

a = "_".join(res)
print("a = " + a)
b = "".join(chr(int(c, 36) + 10) for c in a.split("_"))
if u == b:
    print("Correct")
else:
    print("Incorrect")
