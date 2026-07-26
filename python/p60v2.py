from readline import append_history_file

from euler_lib import SieveOfEratosthenes, is_prime_by_division


def is_prime(val, sieve):
    if val < sieve.size:
        return sieve.is_prime(val)
    return is_prime_by_division(val)


def test_pair_update_cache(p, sieve, cache):
    if p in cache:
        return cache[p]
    v1 = int(str(p[0]) + str(p[1]))
    v2 = int(str(p[1]) + str(p[0]))
    result = is_prime(v1, sieve) and is_prime(v2, sieve)
    cache[p] = result
    return result


def can_add_to_set(val, a_set, sieve, cache):
    for v in a_set:
        if not test_pair_update_cache((v, val), sieve, cache):
            return False
    return True


print("computing prime")
sieve = SieveOfEratosthenes(10000)
primes = sieve.get_primes_in_sieve()
print("number of primes", len(primes))
size = 4
# size = 5
cache = {}
candidates = []

for i in range(len(primes) - 1):
    this_set = set()
    this_set.add(primes[i])
    for j in range(i + 1, len(primes)):
        if can_add_to_set(primes[j], this_set, sieve, cache):
            this_set.add(primes[j])
    if len(this_set) >= size:
        candidates.append(this_set)

print("number of candidate sets", len(candidates))

min_sum = None
min_combo = None
for c in candidates:
    this_sum = sum(c)
    if not min_sum or this_sum < min_sum:
        min_sum = this_sum
        min_combo = c
print(min_sum, min_combo)
