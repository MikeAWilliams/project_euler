import itertools

from euler_lib import from_digits

digits = []
for i in range(10):
    digits.append(i)

target_devisors = [2,3,5,7,11,13,17]
perm_iterator = itertools.permutations(digits)
valid_pd = []
for p in perm_iterator:
    lp = list(p)
    valid = True
    for target_index in range(len(target_devisors)):
       three_digits = lp[target_index + 1: target_index + 4]
       three_num = from_digits(three_digits)
       if three_num % target_devisors[target_index] != 0:
           valid = False
           break
    if valid:
       valid_pd.append(from_digits(lp))
print(valid_pd)
sum = 0
for v in valid_pd:
    sum += v
print(sum)
