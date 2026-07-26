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


def search_for_the_and_for():
    for l1 in string.ascii_lowercase:
        for l2 in string.ascii_lowercase:
            for l3 in string.ascii_lowercase:
                key = l1 + l2 + l3
                decoded = decrypt(key, coded_ints)
                if "the" in decoded and "for" in decoded:
                    print()
                    print(key, decoded)
                    print()


coded_ints = read_file()

# search for spaces
best_key = ""
best_count = 0
for l1 in string.ascii_lowercase:
    for l2 in string.ascii_lowercase:
        for l3 in string.ascii_lowercase:
            key = l1 + l2 + l3
            decoded = decrypt(key, coded_ints)
            spaces = decoded.count(" ")
            if spaces > best_count:
                best_key = key
                best_count = spaces
print("best key", best_key)
decoded = decrypt(best_key, coded_ints)
print("message\n", decoded, "\n")
ascii = [ord(c) for c in decoded]
print("answser", sum(ascii))
