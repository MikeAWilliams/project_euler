from euler_lib import SieveOfEratosthenes, get_digits


def get_new_rotations(number):
    result = []
    digits = get_digits(number)
    if len(digits) == 1:
        return result
    digits.reverse()
    for shifts in range(len(digits)-1):
        zero_item = digits[0]
        digits = digits[1:]
        digits.append(zero_item)
        new_num = 0
        for index in range(len(digits)):
            new_num += digits[index] * pow(10, len(digits) - index -  1)
        result.append(new_num)
    return result

size = 1000000
sieve = SieveOfEratosthenes(size)
count = 5
for candidate in range(13, size, 2):
    if sieve.is_prime(candidate):
        rotations = get_new_rotations(candidate)
        valid = True
        for r in rotations:
            if r >= size or not sieve.is_prime(r):
                valid = False
                continue
        if valid:
            count +=1
print(count)
