from itertools import combinations

from euler_lib import SieveOfEratosthenes, is_prime_by_division


def test_combo(combo):
    pairs = combinations(combo, 2)
    for p in pairs:
        v1 = int(str(p[0]) + str(p[1]))
        if not is_prime_by_division(v1):
            return False
        v2 = int(str(p[1]) + str(p[0]))
        if not is_prime_by_division(v2):
            return False
    return True


sieve = SieveOfEratosthenes(1000)
primes = sieve.get_primes_in_sieve()
size = 4
combos = combinations(primes, size)

min_sum = None
min_combo = None
for c in combos:
    if test_combo(c):
        this_sum = sum(c)
        if not min_sum or this_sum < min_sum:
            min_combo = c
            min_sum = this_sum
print(min_sum, min_combo)
