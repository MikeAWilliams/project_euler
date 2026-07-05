import math


def get_digits(n):
    result = []
    while n > 0:
        digit = n % 10
        result.append(digit)
        n = n // 10
    return result

result = 0
for num in range(1000, 10000000):
    digits = get_digits(num)
    sum = 0
    for d in digits:
        sum += math.pow(d,5)
    if sum == num:
        result+=num
        print("num", num, "current sum", result)
print("result", result)
