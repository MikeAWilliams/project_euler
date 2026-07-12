from euler_lib import SieveOfEratosthenes

size = 1000000

print("computing sieve")
sieve = SieveOfEratosthenes(size)
primes = sieve.get_primes_in_sieve()

print("searching")
max_len = 0
max_prime = 0
for i in range(len(primes)-1):
    if i % 1000 == 0:
        print(i, max_len, max_prime)
    sum = primes[i]
    for j in range(i+1, len(primes)):
        sum += primes[j]
        if sum >= size:
            break
        if sieve.is_prime(sum):
            consecutive = j - i + 1
            if consecutive > max_len:
                max_len = consecutive
                max_prime = sum
print(max_len, max_prime)
