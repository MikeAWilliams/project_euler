import itertools

from euler_lib import SieveOfEratosthenes, from_digits


def get_digits_one_to_n(n):
    digits = []
    for i in range(1,n+1):
       digits.append(i)
    return digits


print("computing sieve")
sieve = SieveOfEratosthenes(987654322)

print("searching")
largets_prime = 0
for n in range(1, 10):
    digits = get_digits_one_to_n(n)
    perm_iterator = itertools.permutations(digits)
    for p in perm_iterator:
        num = from_digits(list(p))
        if sieve.is_prime(num):
            if num > largets_prime:
                largets_prime = num
print(largets_prime)
