import struct
from typing import List

class F2_32:
    def __init__(self, val: int):
        self.val = val & 0xffffffff
    def __add__(self, other):
        return F2_32(self.val + other.val)
    def __sub__(self, other):
        return F2_32(self.val - other.val + 0xffffffff + 1)
    def __xor__(self, other):
        return F2_32(self.val ^ other.val)
    def __lshift__(self, nbit: int):
        left = (self.val << nbit) & 0xffffffff
        right = (self.val & 0xffffffff) >> (32 - nbit)
        return F2_32(left | right)
    def __rshift__(self, nbit: int):
        left = (self.val & 0xffffffff) >> nbit
        right = (self.val << (32 - nbit)) & 0xffffffff
        return F2_32(left | right)
    def __repr__(self):
        return hex(self.val)
    def __int__(self):
        return int(self.val)
    
def serialize(state: List[F2_32]) -> List[bytes]:
    return b''.join([ struct.pack('<I', int(s)) for s in state ])