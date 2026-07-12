pentagonals = []
for n in range(1, 10000):
    pentagonals.append(n*(3*n-1)//2)
p_set = set(pentagonals)

for i1 in range(len(pentagonals)-1):
    for i2 in range(i1+1, len(pentagonals)):
        sum = pentagonals[i2]+pentagonals[i1]
        dif = pentagonals[i2]-pentagonals[i1]
        if sum in p_set and dif in p_set:
            print(i1, i2, sum, dif)
