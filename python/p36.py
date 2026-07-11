from euler_lib import get_digits


def is_binary_palindrome(decimal_num):
    digits = bin(decimal_num)
    digits = digits[2:]
    if len(digits) == 1:
        return True
    for index in range(len(digits) // 2):
        if digits[index] != digits[len(digits)-index-1]:
            return False
    return True


def is_decimal_palindrome(num):
    digits = get_digits(num)
    if len(digits) == 1:
        return True
    for index in range(len(digits) // 2):
        if digits[index] != digits[len(digits)-index-1]:
            return False
    return True

def is_double_palindrome(num):
    return is_decimal_palindrome(num) and is_binary_palindrome(num)


sum = 0
for i in range(1000000):
    if is_double_palindrome(i):
        sum += i
print(sum)
