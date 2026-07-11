def dumb_pow(base, pow):
    result = 1
    for i in range(pow):
        result *= base
    return result

max = 1000
sum = 0
for i in range(1,max + 1):
    sum += dumb_pow(i, i)

# manual get the last 10 digits
digits = []
for i in range(10):
    digit = sum % 10
    digits.append(digit)
    sum = sum // 10
digits.reverse()
print(digits)
