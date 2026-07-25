from euler_lib import from_digits, get_digits, is_decimal_palindrome


def reverse_add(num):
    digits = get_digits(num)
    # digits are already reversed
    return num + from_digits(digits)


def is_lychrel50(num):
    for i in range(50):
        num = reverse_add(num)
        if is_decimal_palindrome(num):
            return False
    return True


print("reverse add test", reverse_add(349))
print("349 lychrel test", is_lychrel50(349))
print("10677 lycrel test", is_lychrel50(10677))

count = 0
for num in range(10000):
    if is_lychrel50(num):
        count += 1

print(count)
