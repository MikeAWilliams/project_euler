coins = [ 200, 100, 50, 20, 10, 5, 2, 1 ]

def get_valid_counts_of_coin(remainder, index):
    result = []
    if index >= len(coins):
        return result
    multiple = 1
    while multiple*coins[index] <= remainder:
        result.append(multiple)
        multiple += 1
    return result

def recursive_ways(amount, index):
    total = 0
    for count in get_valid_counts_of_coin(amount, index):
        total += recursive_ways(amount - count*coins[index], index + 1)
    return total

print(get_valid_counts_of_coin(200, 0))
print(recursive_ways(200, 0))
