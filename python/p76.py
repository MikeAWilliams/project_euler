# compute the number of ways to count to target using numbers < target
#target = 100
target = 5

# to make the counting work assume there is one way to make 0 which is an empty set
ways = {
    0:1
}
# start without knowing how many ways to make any given target > 0
for n in range(1,target+1):
    ways[n] = 0

# here a largest_coin means a number < target that I am using to add up
# so when largets_coin is 1 I can only use 1 + 1+ 1 until we get to target. that 1 way
for largest_coin in range(1,target):
   # update the number of ways I can write n as a sum using this largetst_coin
   for n in range(largest_coin,target + 1):
      # we already have the ways for the last largest_coin
      # so add to that the ways to make n - largest
      ways[n] += ways[n - largest_coin]

   print(largest_coin, ways)

print(ways[target])
