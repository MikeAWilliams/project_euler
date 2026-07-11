
from euler_lib import SieveOfEratosthenes

sieve_size = 1000000
consecutive_target = 4

sieve = SieveOfEratosthenes(sieve_size)
primes = sieve.get_primes_in_sieve()

results = []
for candidate in range(1, sieve_size):
    factors = sieve.get_prime_factors(candidate, primes)
    distince_factor = set(factors)
    if len(distince_factor) == consecutive_target:
        results.append(candidate)
        if len(results) == consecutive_target:
            break
    else:
        results = []
print(results)
