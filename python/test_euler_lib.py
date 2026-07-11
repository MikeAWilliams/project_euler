from euler_lib import get_digits, get_factors


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
