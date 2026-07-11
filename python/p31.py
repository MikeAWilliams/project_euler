coins = [ 200, 100, 50, 20, 10, 5, 2, 1 ]

def get_valid_counts_of_coin(remainder, index):
    result = []
    if index >= len(coins):
        return result
    # need to include the zero here so that the recursion can contine with zero of this coin
    multiple = 0
    while multiple*coins[index] <= remainder:
        result.append(multiple)
        multiple += 1
    return result

def recursive_ways(remainder, index):
    # this iteration hit the total already
    if remainder == 0:
        return 1
    total = 0
    for count in get_valid_counts_of_coin(remainder, index):
        total += recursive_ways(remainder - count*coins[index], index + 1)
    return total
print(recursive_ways(200, 0))
