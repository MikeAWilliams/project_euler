import itertools

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
    for number_to_replace in range(1, len(digits)):
        replace_indexes = itertools.combinations(range(len(digits)), number_to_replace)
        for new_value in range(10):
            mutate_digits = digits.copy()
            for a_replace_set in replace_indexes:
                for a_replace in a_replace_set:
                    mutate_digits[a_replace] = new_value
            new_num = from_digits(mutate_digits)
            if sieve.is_prime(new_num):
                family.append(new_num)
