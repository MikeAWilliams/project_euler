from euler_lib import SieveOfEratosthenes, from_digits, get_digits

size = 1000000

print("computing sieve")
sieve = SieveOfEratosthenes(size)
primes = sieve.get_primes_in_sieve()

print("calculating")
for p in primes:
    digits = get_digits(p)
    digits.reverse()
    family = [p]
    for number_to_replace in range(len(digits)):
        replace_indexes = itertools.combinations(range(len(digits)), number_to_replace)
        print(replace_indexes)
        exit(1)
        for new_value in range(10):
            for replace_index in range(len(digits)):
