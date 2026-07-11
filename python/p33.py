from euler_lib import get_digits, get_factors

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
        # only class 3 showed up so I'm only coding for that one
        #if num_digits[0] == den_digits[0]:
        #    fake_divide = num_digits[1] / den_digits[1]
        #    if fake_divide == real_divide:
        #        print("case 1", num, den)

        #if num_digits[1] != 0 and num_digits[1] == den_digits[1]:
        #    fake_divide = num_digits[0] / den_digits[0]
        #    if fake_divide == real_divide:
        #        print("case 2", num, den)

        if num_digits[0] == den_digits[1]:
            fake_divide = num_digits[1] / den_digits[0]
            if fake_divide == real_divide:
                print("case 3", num, den)
                result_num *= num
                result_den *= den

        # only class 3 showed up so only coding for that one
        #if num_digits[1] == den_digits[0]:
        #    fake_divide = num_digits[0] / den_digits[1]
        #    if fake_divide == real_divide:
        #        print("case 4", num, den)
print(result_num, result_den)
num_factors = get_factors(result_num)
den_factors = get_factors(result_den)
#print(num_factors)
#print(den_factors)
common_factors = []
for num_fac in num_factors[1:]:
    for den_fac in den_factors[1:]:
        if num_fac == den_fac:
            common_factors.append(num_fac)

# definitely being dumb here but also lazy
reduced_num = result_num
reduced_den = result_den
for fac in common_factors:
    if reduced_num % fac == 0 and reduced_den % fac == 0:
        reduced_num /= fac
        reduced_den /= fac
print(reduced_num, reduced_den)
