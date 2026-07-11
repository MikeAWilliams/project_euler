from euler_lib import SieveOfEratosthenes

sieve = SieveOfEratosthenes(1000)
primes = sieve.get_primes_in_sieve()
print(primes)

for candidate in range(9, 1000, 2):
   if sieve.is_prime(candidate):
      continue
