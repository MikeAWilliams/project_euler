def get_next(n):
    n_str = str(n)
    sum = 0
    for d in n_str:
        d_int = int(d)
        sum += d_int * d_int
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
