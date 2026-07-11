coins = [ 200, 100, 50, 20, 10, 5, 2, 1 ]

# plan to place tupls of coin value and count
results = []

while len(coins) > 0:
    top_coin = coins[0]
    coins = coins[1:]
    starter_list = [(top_coin,1)]
    multiple = 2
    while multiple * top_coin <= 200:
        starter_list.append((top_coin, multiple))
        multiple += 1
    print(starter_list)
