from euler_lib import SieveOfEratosthenes, get_digits

size = 1000000000

print("computing sieve")
sieve = SieveOfEratosthenes(size)
primes = sieve.get_primes_in_sieve()

print("calculating")
for p in primes:
    digits = get_digits(p)
    if len(digits) < 9:
        continue
    print(p)
