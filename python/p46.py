from euler_lib import SieveOfEratosthenes

sieve = SieveOfEratosthenes(1000)
for candidate in range(9, 1000, 2):
   if sieve.is_prime(candidate):
      continue
    print(candidate)
