import math

from euler_lib import get_digits

target_digits = []
for i in range(7):
    digit = int(math.pow(10, i))
    target_digits.append(digit)
print(target_digits)

target_index = 0
digits = []
digit_counter = 0
for i in range(target_digits[len(target_digits)-1] + 1):
    i_digits = get_digits(i)
    new_digit_counter = digit_counter + len(i_digits)
    if new_digit_counter >= target_digits[target_index]:
        i_digits.reverse()
        delta = target_digits[target_index]-digit_counter-1
        print(i, target_index, target_digits[target_index], digit_counter, new_digit_counter, delta)
        digits.append(i_digits[delta])
        target_index += 1
    digit_counter = new_digit_counter
    if target_index >= len(target_digits):
        break
print("digits", digits)
product = 1
for d in digits:
    product *=d
print("final answer", product)
