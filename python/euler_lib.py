import math


def get_factors(n):
    result = [1]
    if n > 1:
        result.append(n)
    stop = int(math.isqrt(n))
    for i in range(2, stop + 1):
        if n % i == 0:
            result.append(i)
            second = n // i
            if second != i:
                result.append(second)
    result.sort()
    return result


def get_digits(n):
    result = []
    while n > 0:
        digit = n % 10
        result.append(digit)
        n = n // 10
    return result
