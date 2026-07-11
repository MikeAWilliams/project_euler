from euler_lib import get_digits

for num in range(10, 100):
    for den in range(10, 100):
        real_divide = num/den
        num_digits = get_digits(num)
        den_digits = get_digits(den)
        if num_digits[0] == den_digits[0]:
            fake_divide = num_digits[1] / den_digits[1]
            if fake_divide == real_divide:
                print(num, den)

        if num_digits[1] == den_digits[1]:
            fake_divide = num_digits[0] / den_digits[0]
            if fake_divide == real_divide:
                print(num, den)
