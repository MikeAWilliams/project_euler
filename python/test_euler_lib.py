from euler_lib import get_digits


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
