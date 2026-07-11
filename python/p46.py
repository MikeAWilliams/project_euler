from euler_lib import SieveOfEratosthenes

max_prime = 10000

def is_valid(candidate, primes):
    for p in primes:
        if p >= candidate:
            return False
        for to_square in range(1, 1000):
            test_val = p + 2 * to_square*to_square
            if test_val == candidate:
                return True
    return False

sieve = SieveOfEratosthenes(max_prime)
primes = sieve.get_primes_in_sieve()

for candidate in range(9, max_prime, 2):
   if sieve.is_prime(candidate):
      continue
   if not is_valid(candidate, primes):
       print(candidate)
       exit(0)
print("failed")
exit(1)
