import itertools

from euler_lib import SieveOfEratosthenes, from_digits, get_digits

size = 1000000

print("computing sieve")
sieve = SieveOfEratosthenes(size)
primes = sieve.get_primes_in_sieve()

print("calculating families")
families = []
for p in primes:
    digits = get_digits(p)
    digits.reverse()
    for number_to_replace in range(1, len(digits)):
        for replace_set in itertools.combinations(range(len(digits)), number_to_replace):
            family = []
            for new_value in range(10):
                mutate_digits = digits.copy()
                for index_to_replace in replace_set:
                    mutate_digits[index_to_replace] = new_value
                if mutate_digits[0] == 0:
                    continue
                new_num = from_digits(mutate_digits)
                if sieve.is_prime(new_num):
                    family.append(new_num)
            if len(family) == 8:
                families.append(family)

print("smallest member")
smallest = 1000000000000000
for f in families:
    for v in f:
        if v < smallest:
            smallest = v
print(smallest)
