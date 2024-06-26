FLAG = b'FAKE{XXXXXXXXXXXXXXXXXXXXXX}'

N = 7
M = 17005450388330379
WORD_SIZE = 32
WORD_MASK = (1 << WORD_SIZE) - 1

def encrypt(m):
    state = [int.from_bytes(m[i:i+4]) for i in range(0, len(m), 4)]
    assert len(state) == N

    def xor_shift():
        nonlocal state
        t = state[0] ^ ((state[0] << 11) & WORD_MASK)
        for i in range(N-1):
            state[i] = state[i+1]
        state[-1] = (state[-1] ^ (state[-1] >> 19)) ^ (t ^ (t >> 8))

    for _ in range(M):
        xor_shift()

    return state

print("N = ", N)
print("M = ", M)
print("WORD_SIZE = ", WORD_SIZE)
print("state = ", encrypt(FLAG))
