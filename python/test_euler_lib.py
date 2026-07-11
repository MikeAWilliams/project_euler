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
