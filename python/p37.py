from euler_lib import SieveOfEratosthenes, get_digits

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


size = 1000000
sieve = SieveOfEratosthenes(size)
count = 0
sum = 0
for candidate in range(9, size, 2):
    if is_valid(candidate, sieve, size)
        count += 1
        sum += candidate
print(count, sum)
