from euler_lib import SieveOfEratosthenes, get_digits, get_factors, from_digits


def test_get_digits_single():
    assert get_digits(5) == [5]


def test_get_digits_multi():
    assert get_digits(123) == [3, 2, 1]


def test_get_digits_trailing_zero():
    assert get_digits(100) == [0, 0, 1]


def test_get_digits_large():
    assert get_digits(9876) == [6, 7, 8, 9]


def test_get_digits_zero():
    assert get_digits(0) == []


def test_from_digits_single():
    assert from_digits([5]) == 5


def test_from_digits_multi():
    assert from_digits([1, 2, 3]) == 123


def test_from_digits_leading_zero():
    assert from_digits([0, 4, 2]) == 42


def test_from_digits_empty():
    assert from_digits([]) == 0


def test_from_digits_round_trip():
    digits = get_digits(9876)
    digits.reverse()
    assert from_digits(digits) == 9876


def test_get_factors_one():
    assert get_factors(1) == [1]

def test_get_factors_prime():
    assert get_factors(3) == [1, 3]

def test_get_factors_perfect_square():
    assert get_factors(9) == [1, 3, 9]

def test_get_factors_composite():
    assert get_factors(6) == [1, 2, 3, 6]

def test_get_factors_28():
    assert get_factors(28) == [1, 2, 4, 7, 14, 28]

def test_get_factors_500():
    assert get_factors(500) == [1, 2, 4, 5, 10, 20, 25, 50, 100, 125, 250, 500]


def test_sieve_primes_under_30():
    sieve = SieveOfEratosthenes(30)
    assert [n for n in range(30) if sieve.is_prime(n)] == [
        2, 3, 5, 7, 11, 13, 17, 19, 23, 29
    ]


def test_sieve_is_prime_true():
    sieve = SieveOfEratosthenes(30)
    assert sieve.is_prime(13)


def test_sieve_is_prime_false():
    sieve = SieveOfEratosthenes(30)
    assert not sieve.is_prime(9)


def test_sieve_one_is_not_prime():
    sieve = SieveOfEratosthenes(30)
    assert not sieve.is_prime(1)


def test_sieve_nth_prime_first():
    sieve = SieveOfEratosthenes(30)
    assert sieve.get_nth_prime(1) == 2


def test_sieve_nth_prime():
    sieve = SieveOfEratosthenes(30)
    assert sieve.get_nth_prime(6) == 13


def test_get_primes_in_sieve_under_30():
    sieve = SieveOfEratosthenes(30)
    assert sieve.get_primes_in_sieve() == [
        2, 3, 5, 7, 11, 13, 17, 19, 23, 29
    ]


def test_get_primes_in_sieve_small():
    sieve = SieveOfEratosthenes(3)
    assert sieve.get_primes_in_sieve() == [2]


def test_get_primes_in_sieve_matches_is_prime():
    sieve = SieveOfEratosthenes(100)
    assert sieve.get_primes_in_sieve() == [
        n for n in range(100) if sieve.is_prime(n)
    ]


def test_get_prime_factors_prime():
    sieve = SieveOfEratosthenes(1000)
    primes = sieve.get_primes_in_sieve()
    assert sieve.get_prime_factors(13, primes) == [13]


def test_get_prime_factors_prime_power():
    sieve = SieveOfEratosthenes(1000)
    primes = sieve.get_primes_in_sieve()
    assert sieve.get_prime_factors(8, primes) == [2, 2, 2]


def test_get_prime_factors_composite():
    sieve = SieveOfEratosthenes(1000)
    primes = sieve.get_primes_in_sieve()
    assert sieve.get_prime_factors(12, primes) == [2, 2, 3]


def test_get_prime_factors_distinct():
    sieve = SieveOfEratosthenes(1000)
    primes = sieve.get_primes_in_sieve()
    assert sieve.get_prime_factors(30, primes) == [2, 3, 5]


def test_get_prime_factors_large_prime_factor():
    # 998 = 2 * 499; 499 is only reached via the sqrt leftover
    sieve = SieveOfEratosthenes(1000)
    primes = sieve.get_primes_in_sieve()
    assert sieve.get_prime_factors(998, primes) == [2, 499]


def test_get_prime_factors_repeated_and_mixed():
    # 360 = 2^3 * 3^2 * 5
    sieve = SieveOfEratosthenes(1000)
    primes = sieve.get_primes_in_sieve()
    assert sieve.get_prime_factors(360, primes) == [2, 2, 2, 3, 3, 5]


def test_get_prime_factors_sorted():
    sieve = SieveOfEratosthenes(1000)
    primes = sieve.get_primes_in_sieve()
    factors = sieve.get_prime_factors(924, primes)
    assert factors == sorted(factors)


def test_get_prime_factors_product_reconstructs():
    sieve = SieveOfEratosthenes(1000)
    primes = sieve.get_primes_in_sieve()
    for n in range(2, 1000):
        factors = sieve.get_prime_factors(n, primes)
        product = 1
        for f in factors:
            product *= f
        assert product == n
        assert all(sieve.is_prime(f) for f in factors)
