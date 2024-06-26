from mycipher import MyCipher
import hashlib
import datetime
import random

isLogged = False
current_user = ''
d = {}

def make_token(data1: str, data2: str):
    sha256 = hashlib.sha256()
    sha256.update(data1.encode())
    right = sha256.hexdigest()[:20]
    sha256.update(data2.encode())
    left = sha256.hexdigest()[:12]
    token = left + right
    return token

def main():
    print('Welcome to the super secure encryption service!')
    while True:
        print('Select an option:')
        print('1. Register')
        print('2. Login')
        print('3. Logout')
        print('4. Encrypt')
        print('5. Exit')
        choice = input('> ')
        if choice == '1':
            Register()
        elif choice == '2':
            Login()
        elif choice == '3':
            Logout()
        elif choice == '4':
            Encrypt()
        elif choice == '5':
            print('Goodbye!')
            break
        else:
            print('Invalid choice')

def Register():
    global d
    username = input('Enter username: ')
    if username in d:
        print('Username already exists')
        return
    dt_now = datetime.datetime.now()
    minutes = dt_now.minute
    sec = dt_now.second
    data1 = f'user: {username}, {minutes}:{sec}'
    data2 = f'{username}'+str(random.randint(0, 10))
    d[username] = make_token(data1, data2)
    print('Registered successfully!')
    print('Your token is:', d[username])
    return

def Login():
    global isLogged
    global d
    global current_user
    username = input('Enter username: ')
    if username not in d:
        print('Username does not exist')
        return
    token = input('Enter token: ')
    if d[username] != token:
        print('Invalid token')
        return
    isLogged = True
    current_user = username
    print(f'Logged in successfully! Hi {username}!')
    return

def Logout():
    global isLogged
    global current_user
    isLogged = False
    current_user = ''
    print('Logged out successfully!')
    return

def Encrypt():
    global isLogged
    global current_user
    if not isLogged:
        print('You need to login first')
        return
    token = d[current_user]
    sha256 = hashlib.sha256()
    sha256.update(token.encode())
    key = sha256.hexdigest()[:32]
    nonce = token[:12]
    cipher = MyCipher(key.encode(), nonce.encode())
    plaintext = input('Enter plaintext: ')
    ciphertext = cipher.encrypt(plaintext.encode())
    print('username:', current_user)
    print('Ciphertext:', ciphertext.hex())
    return

if __name__ == '__main__':
    main()