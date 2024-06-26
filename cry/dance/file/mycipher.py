from utils import *

class MyCipher:
    def __init__(self, key: bytes, nonce: bytes):
        self.key = key
        self.nonce = nonce
        self.counter = 1
        self.state = List[F2_32]

    def __quarter_round(self, a: F2_32, b: F2_32, c: F2_32, d: F2_32):
        a += b; d ^= a; d <<= 16
        c += d; b ^= c; b <<= 12
        a += b; d ^= a; d <<= 8
        c += d; b ^= c; b <<= 7
        return a, b, c, d
    
    def __Qround(self, idx1, idx2, idx3, idx4):
        self.state[idx1], self.state[idx2], self.state[idx3], self.state[idx4] = \
            self.__quarter_round(self.state[idx1], self.state[idx2], self.state[idx3], self.state[idx4])

    def __update_state(self):
        for _ in range(10):
            self.__Qround(0, 4, 8, 12)
            self.__Qround(1, 5, 9, 13)
            self.__Qround(2, 6, 10, 14)
            self.__Qround(3, 7, 11, 15)
            self.__Qround(0, 5, 10, 15)
            self.__Qround(1, 6, 11, 12)
            self.__Qround(2, 7, 8, 13)
            self.__Qround(3, 4, 9, 14)

    def __get_key_stream(self, key: bytes, counter: int, nonce: bytes) -> bytes:
        constants = [F2_32(x) for x in struct.unpack('<IIII', b'expand 32-byte k')]
        key = [F2_32(x) for x in struct.unpack('<IIIIIIII', key)]
        counter = [F2_32(counter)]
        nonce = [F2_32(x) for x in struct.unpack('<III', nonce)]
        self.state = constants + key + counter + nonce
        initial_state = self.state[:]
        self.__update_state()
        self.state = [x + y for x, y in zip(self.state, initial_state)]
        return serialize(self.state)
    
    def __xor(self, a: bytes, b: bytes) -> bytes:
        return bytes([x ^ y for x, y in zip(a, b)])

    def encrypt(self, plaintext: bytes) -> bytes:
        encrypted_message = bytearray(0)

        for i in range(len(plaintext)//64):
            key_stream = self.__get_key_stream(self.key, self.counter + i, self.nonce)
            encrypted_message += self.__xor(plaintext[i*64:(i+1)*64], key_stream)

        if len(plaintext) % 64 != 0:
            key_stream = self.__get_key_stream(self.key, self.counter + len(plaintext)//64, self.nonce)
            encrypted_message += self.__xor(plaintext[(len(plaintext)//64)*64:], key_stream[:len(plaintext) % 64])

        return bytes(encrypted_message)
