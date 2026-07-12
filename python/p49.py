from euler_lib import SieveOfEratosthenes, get_digits


def all_three_are_permutations(c1, c2, c3):
   d1 = get_digits(c1)
   d2 = get_digits(c2)
   d3 = get_digits(c3)
   if len(d1) != len(d2) or len(d1) != len(d3):
      return False
   for d in d1:
      found_in_2 = False
      for d2_index in range(len(d2)):
         if d2[d2_index] == d:
            found_in_2 = True
            break
      if not found_in_2:
         return False
      d2.remove(d)
      found_in_3 = False
      for d3_index in range(len(d3)):
         if d3[d3_index] == d:
             found_in_3 = True
             break
      if not found_in_3:
          return False
      d3.remove(d)
   return True

sieve = SieveOfEratosthenes(100000)

for candidate in range(10001 - 3330*2):
    if not sieve.is_prime(candidate):
        continue
    candidate1 = candidate + 3330
    if not sieve.is_prime(candidate1):
        continue
    candidate2 = candidate1 + 3330
    if not sieve.is_prime(candidate2):
        continue
    if all_three_are_permutations(candidate, candidate1, candidate2):
        print(candidate, candidate1, candidate2)
