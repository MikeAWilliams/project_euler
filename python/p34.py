from math import factorial

from euler_lib import get_digits


def precompute_digit_factorial():
    result = {}
    for i in range(10):
        result[i] = factorial(i)
    return result

def test_number(number, factorial_cache):
    digits = get_digits(number)
    if len(digits) <= 1:
        return False
    result = 0
    for d in digits:
        result += factorial_cache[d]
    return result == number

sum = 0
factorial_cache = precompute_digit_factorial()
for number in range(10000000):
    if test_number(number, factorial_cache):
        print(number)
        sum+=number
print(sum)
