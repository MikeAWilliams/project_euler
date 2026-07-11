from euler_lib import get_digits


def is_pandigital9(n):
    digits = get_digits(n)
    if len(digits) != 9:
        return False
    digits_set = set(digits)
    if len(digits_set) != 9:
        return False
    if 0 in digits_set:
       return False
    return True

largest = 0
largest_primary = 0
largest_n = 0
for primary in range(1,99999):
    result = ""
    for n in range(1,10):
       product = primary * n
       result += str(product)
       result_num = int(result)
       if is_pandigital9(result_num):
           print(primary, n, result_num)
           if result_num > largest:
               largest = result_num
               largest_primary = primary
               largest_n = n
print("final", largest_primary, largest_n, largest)
