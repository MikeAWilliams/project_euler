from math import gcd

from euler_lib import get_digits

result_num = 1
result_den = 1
for num in range(10, 100):
    for den in range(10, 100):
        if num >= den:
            continue
        real_divide = num/den
        num_digits = get_digits(num)
        den_digits = get_digits(den)
        if num == den or den_digits[0] == 0:
            continue

        # this is the only case that resulted in matches.
        if num_digits[0] == den_digits[1]:
            fake_divide = num_digits[1] / den_digits[0]
            if fake_divide == real_divide:
                print("case 3", num, den)
                result_num *= num
                result_den *= den

print(result_num, result_den)
factor = gcd(result_num, result_den)
while factor != 1:
    result_num = result_num // factor
    result_den = result_den // factor
    factor = gcd(result_num, result_den)
print(result_num, result_den)
