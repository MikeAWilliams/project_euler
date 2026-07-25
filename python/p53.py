import math

count = 0
for n in range(1,101):
    for r in range(1,n+1):
        choose = math.comb(n,r)
        if choose > 1000000:
            count+=1
print(count)
