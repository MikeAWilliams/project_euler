
from euler_lib import SieveOfEratosthenes

size_size = 10000

sieve = SieveOfEratosthenes(size_size)
primes = sieve.get_primes_in_sieve()

print(sieve.get_prime_factors(646, primes))
