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
    # replace between 1 and all but one of the digits
    for number_to_replace in range(1, len(digits)):
        # compute all the indicies to replace for this number of replacements
        for replace_set in itertools.combinations(range(len(digits)), number_to_replace):
            # a family is all the primes after these digits are replaced with any values 0-9
            family = []
            for new_value in range(10):
                mutate_digits = digits.copy()
                # replace all the values at each replace index in the mutate_digits with new_value
                for index_to_replace in replace_set:
                    mutate_digits[index_to_replace] = new_value
                # leading zeros make the number smaller and invalid for the family
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
