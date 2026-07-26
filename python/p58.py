from euler_lib import SieveOfEratosthenes

print("computing the sieve")
grid_size = 30001
sieve = SieveOfEratosthenes(grid_size * grid_size)
print("calculating")
# unfortunately building the actual grid is to memory expensive.
# Instead we will compute the new corner values at each size step

prime_count = 0
diagonal_count = 1
for size in range(3, grid_size, 2):
    size_minus_1 = size - 1
    c1 = size * size
    c2 = c1 - size_minus_1
    c3 = c1 - 2 * size_minus_1
    c4 = c1 - 3 * size_minus_1
    diagonal_count += 4
    if sieve.is_prime(c1):
        prime_count += 1
    if sieve.is_prime(c2):
        prime_count += 1
    if sieve.is_prime(c3):
        prime_count += 1
    if sieve.is_prime(c4):
        prime_count += 1
    ratio = prime_count / diagonal_count
    if ratio < 0.1:
        print(size)
        break
