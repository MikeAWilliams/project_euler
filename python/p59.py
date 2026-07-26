import csv
import string


def read_file():
    with open("0059_cipher.txt") as file:
        reader = csv.reader(file)
        return [int(x) for x in next(reader)]


def decrypt(key, cipher_message):
    ascii_key = [ord(char) for char in key]
    decoded_message = ""
    for i, c in enumerate(cipher_message):
        key_index = i % len(key)
        decoded = c ^ ascii_key[key_index]
        decoded_message += chr(decoded)
    return decoded_message


coded_ints = read_file()
for l1 in string.ascii_lowercase:
    for l2 in string.ascii_lowercase:
        for l3 in string.ascii_lowercase:
            key = l1 + l2 + l3
            decoded = decrypt(key, coded_ints)
            if "the" in decoded and "for" in decoded:
                print()
                print(key, decoded)
                print()

print()
print()
print()
print("using one of the above")
key = "exp"
decoded = decrypt(key, coded_ints)
ascii = [ord(c) for c in decoded]
print(sum(ascii))
