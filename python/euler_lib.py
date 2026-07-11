import math
from argparse import ArgumentError


def get_factors(n):
    result = [1]
    if n > 1:
        result.append(n)
    stop = int(math.isqrt(n))
    for i in range(2, stop + 1):
        if n % i == 0:
            result.append(i)
            second = n // i
            if second != i:
                result.append(second)
    result.sort()
    return result


def get_digits(n):
    result = []
    while n > 0:
        digit = n % 10
        result.append(digit)
        n = n // 10
    return result

class SieveOfEratosthenes:
    def __init__(self, size):
        if size <= 2:
            raise Exception(f"Invalide size {size}")
        self.flags = []
        self.size = size
        self.init_flags()
        self.run_sieve()

    def init_flags(self):
        for i in range(self.size):
            self.flags.append(False)
        self.flags[2] = True
        # odd numbers might be prime at first
        for i in range(3, self.size, 2):
            self.flags[i] = True

    def run_sieve(self):
        candidate = 3
        # search up to sqrt(size)
        while candidate*candidate < self.size:
            # starting with 9 (3*3) because even is already false and odd before 9 is prime
            not_prime = candidate * candidate
            # advance forward by 2*candidate to skip even values which are already false
            advance = candidate * 2
            while not_prime < self.size:
                self.flags[not_prime] = False
                not_prime += advance
            # skip ahead to the next candidate that is still flagged prime
            candidate += 2
            while candidate < self.size and not self.flags[candidate]:
                candidate += 2


    def is_prime(self, num):
        if num >= self.size or num < 0:
            raise Exception(f"Invalid number {num}, size {self.size}")
        return self.flags[num]

    def get_nth_prime(self, n):
        if n < 1:
            raise Exception(f"invalid n for nth prime {n}")
        if n == 1:
            return 2
        counter = 1
        for index in range(3, self.size, 2):
            if self.flags[index]:
                counter += 1
                if counter == n:
                    return index
        raise Exception(f"Sieve of size {self.size} does not conatin {n} primes")
