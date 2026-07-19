target = 100
#target = 5
ways = {
    0:1
}
for n in range(1,target+1):
    ways[n] = 0

for largest_coin in range(1,target):
   for n in range(largest_coin,target + 1):
      ways[n] += ways[n - largest_coin]

   #print(largest_coin, ways)

#print(ways)
print(ways[target])
