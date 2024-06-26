from hashlib import sha256
import os
from secrets import randbelow
from secret import flag, cheat_code
import re

challenge_times = 100
hash_strength = int(os.environ.get("HASH_STRENGTH", 10000))

def super_strong_hash(s: str) -> bytes:
    sb = s.encode()
    for _ in range(hash_strength):
        sb = sha256(sb).digest()
    return sb

cheat_code_hash = super_strong_hash(cheat_code)
print(f"hash of cheat code: {cheat_code_hash.hex()}")
print("If you know the cheat code, you will always be accepted!")

secret_number = randbelow(10**10)
secret_code = f"{secret_number:010d}"
print(f"Find the secret code of 10 digits in {challenge_times} challenges!")

def check_code(given_secret_code, given_cheat_code):
    def check_cheat_code(given_cheat_code):
        return super_strong_hash(given_cheat_code) == cheat_code_hash

    digit_is_correct = []
    for i in range(10):
        digit_is_correct.append(given_secret_code[i] == secret_code[i] or check_cheat_code(given_cheat_code))
    return all(digit_is_correct)

given_cheat_code = input("Enter the cheat code: ")
if len(given_cheat_code) > 50:
    print("Too long!")
    exit(1)
for i in range(challenge_times):
    print(f"=====Challenge {i+1:03d}=====")
    given_secret_code = input("Enter the secret code: ")
    if not re.match(r"^\d{10}$", given_secret_code):
        print("Wrong format!")
        exit(1)
    if check_code(given_secret_code, given_cheat_code):
        print("Correct!")
        print(flag)
        exit(0)
    else:
        print("Wrong!")
print("Game over!")
