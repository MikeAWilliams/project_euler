# initial time 13.7s
# change get_next to use % and / 7.97


def get_next(n):
    sum = 0
    while n > 0:
        d = n % 10
        sum += d * d
        n = n // 10
    return sum


def ends89(n):
    while n != 1 and n != 89:
        next = get_next(n)
        n = next
    return n == 89


count = 0
for i in range(1, 10000000):
    if ends89(i):
        count += 1
print(count)
if count != 8581146:
    print("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
