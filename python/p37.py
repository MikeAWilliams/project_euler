from euler_lib import SieveOfEratosthenes, from_digits, get_digits


def is_valid(candidate, sieve, size):
    if candidate >= size:
        return False
    if not sieve.is_prime(candidate):
        return False
    digits = get_digits(candidate)
    digits.reverse()
    # check left to right
    new_digits = digits[1:]
    while len(new_digits) > 0:
        new_num = from_digits(new_digits)
        if not sieve.is_prime(new_num):
            return False
        new_digits = new_digits[1:]
    # check right to left
    new_digits = digits[:]
    new_digits.pop()
    while len(new_digits) > 0:
        new_num = from_digits(new_digits)
        if not sieve.is_prime(new_num):
            return False
        new_digits.pop()
    return True



size = 1000000
sieve = SieveOfEratosthenes(size)

print(is_valid(3797, sieve, size))

count = 0
sum = 0
for candidate in range(9, size, 2):
    if is_valid(candidate, sieve, size):
        count += 1
        sum += candidate
print(count, sum)
