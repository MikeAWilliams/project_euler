def get_digits(n):
    result = []
    while n > 0:
        digit = n % 10
        result.append(digit)
        n = n // 10
    return result
