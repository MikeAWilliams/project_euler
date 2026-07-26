from euler_lib import is_prime_by_division

grid_size = 30001
print("calculating")

prime_count = 0
diagonal_count = 1
for size in range(3, grid_size, 2):
    size_minus_1 = size - 1
    c1 = size * size
    c2 = c1 - size_minus_1
    c3 = c1 - 2 * size_minus_1
    c4 = c1 - 3 * size_minus_1
    diagonal_count += 4
    if is_prime_by_division(c1):
        prime_count += 1
    if is_prime_by_division(c2):
        prime_count += 1
    if is_prime_by_division(c3):
        prime_count += 1
    if is_prime_by_division(c4):
        prime_count += 1
    ratio = prime_count / diagonal_count
    if ratio < 0.1:
        print(size)
        break
