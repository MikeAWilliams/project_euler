from euler_lib import get_digits

max_sum = 0
for a in range(1, 100):
    for b in range(1, 100):
        val = a**b
        digits = get_digits(val)
        this_sum = sum(digits)
        max_sum = max(this_sum, max_sum)
print(max_sum)
