from fractions import Fraction

from euler_lib import get_digits

value = Fraction(1, 1) + Fraction(1, 2)
print(value.numerator, value.denominator)
value = Fraction(1, 1) + Fraction(1, 2 + Fraction(1, 2))
print(value.numerator, value.denominator)
value = Fraction(1, 1) + Fraction(1, 2 + Fraction(1, 2 + Fraction(1, 2)))
print(value.numerator, value.denominator)
value = Fraction(1, 1) + Fraction(
    1, 2 + Fraction(1, 2 + Fraction(1, 2 + Fraction(1, 2)))
)
print(value.numerator, value.denominator)

count = 0
for i in range(1000):
    term = Fraction(1, 2)
    for j in range(i):
        term = Fraction(1, 2 + term)
    value = Fraction(1, 1) + term
    if len(get_digits(value.numerator)) > len(get_digits(value.denominator)):
        count += 1
print("result", count)
