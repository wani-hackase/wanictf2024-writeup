import sys

sys.setrecursionlimit(10000000)

# obfuscate

t = input("Enter the flag: ")
p = lambda x: "".join(chr(ord(c) + 12) for c in x)
s = "".join(chr(ord(c) - 3) for c in p(t))
q = lambda x: "".join(chr(123 ^ ord(c)) for c in x)
u = "".join(q(c) for c in s)

a = "16_10_13_x_6t_4_1o_9_1j_7_9_1j_1o_3_6_c_1o_6r"
b = "".join(chr(int(c, 36) + 10) for c in a.split("_"))

eq = lambda x, y: x == y
res = print("Correct FLAG!" if eq(u, b) else "Incorrect")
