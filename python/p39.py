from math import sqrt

count = {}
for a in range(1, 1000):
    for b in range(a, 1000):
        raw_c = sqrt(a*a + b*b)
        c = int(raw_c)
        if raw_c != c:
            continue
        p = a + b + c
        if p < 1000:
            if p in count:
                count[p]+=1
            else:
                count[p] = 1
print(count[120])
max_len = max(count, key=lambda p: count[p])
print(max_len)
