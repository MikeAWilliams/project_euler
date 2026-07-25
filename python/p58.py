from euler_lib import SieveOfEratosthenes


def get_grid(size):
    return [[0 for _ in range(size)] for _ in range(size)]


def populate_grid(grid):
    i = len(grid) - 1
    j = len(grid) - 1
    side = len(grid)
    v = len(grid) * len(grid)
    dir = "l"
    steps = side - 1
    ring_steps = steps
    while v > 0:
        grid[i][j] = v
        v -= 1
        if dir == "l":
            if steps > 0:
                i -= 1
                steps -= 1
            else:
                dir = "u"
                steps = ring_steps
        if dir == "u":
            if steps > 0:
                j -= 1
                steps -= 1
            else:
                dir = "r"
                steps = ring_steps
        if dir == "r":
            if steps > 0:
                i += 1
                steps -= 1
            else:
                dir = "d"
                steps = ring_steps
        if dir == "d":
            if steps > 1:
                j += 1
                steps -= 1
            else:
                dir = "l"
                ring_steps -= 2
                steps = ring_steps
                i -= 1


print("computing the sieve")
grid_size = 10001
sieve = SieveOfEratosthenes(grid_size * grid_size)
print("solving")
grid = get_grid(grid_size)
populate_grid(grid)
for size in range(9, grid_size, 2):
    origin = (grid_size - size) // 2
    prime_count = 0
    diagonal_count = 0
    for i in range(size):
        diagonal_count += 2
        if sieve.is_prime(grid[origin + i][origin + i]):
            prime_count += 1
        if sieve.is_prime(grid[origin + i][origin + size - i - 1]):
            prime_count += 1
    # above counts the center square twice but since 1 is not prime it won't count primes twice
    diagonal_count -= 1
    fraction = prime_count / diagonal_count
    if fraction < 0.1:
        print(size)
        break
