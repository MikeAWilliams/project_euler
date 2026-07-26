# this algorithm seems obvious and is brute force.
# however it will take around 47d 22h 45m to finish
# so I need to try another approach
import math
import time
from itertools import combinations

from euler_lib import SieveOfEratosthenes, is_prime_by_division


def test_pair_update_cache(p, cache):
    if p in cache:
        return cache[p]
    v1 = int(str(p[0]) + str(p[1]))
    v2 = int(str(p[1]) + str(p[0]))
    result = is_prime_by_division(v1) and is_prime_by_division(v2)
    cache[p] = result
    return result


def test_combo(combo, cache):
    pairs = combinations(combo, 2)
    for p in pairs:
        if not test_pair_update_cache(p, cache):
            return False
    return True


# cache = {}
# print(test_combo([3, 7, 109, 673], cache))
# print(sum([3, 7, 109, 673]))
# exit(1)

print("computing prime")
sieve = SieveOfEratosthenes(10000)
primes = sieve.get_primes_in_sieve()
print("number of primes", len(primes))
# size = 4
size = 5
print("computing combos")
combos = combinations(primes, size)
comb_count = math.comb(len(primes), size)
print(comb_count)


def fmt_secs(s):
    d, rem = divmod(int(s), 86400)
    h, rem = divmod(rem, 3600)
    m, sec = divmod(rem, 60)
    return f"{d}d{h:02d}h{m:02d}m{sec:02d}s"


print("searching")
min_sum = None
min_combo = None
cache = {}
start = time.time()
for i, c in enumerate(combos):
    if i % 100000000 == 0:
        frac = i / comb_count
        elapsed = time.time() - start
        remaining = elapsed / frac - elapsed if frac > 0 else float("inf")
        print(
            i,
            frac,
            "elapsed",
            fmt_secs(elapsed),
            "remaining",
            fmt_secs(remaining) if frac > 0 else "?",
        )
        # estimated time to complete by this is 47d22h45m21s
    if test_combo(c, cache):
        this_sum = sum(c)
        if not min_sum or this_sum < min_sum:
            min_combo = c
            min_sum = this_sum
            print("working value", min_sum, min_combo)
print("final value", min_sum, min_combo)
