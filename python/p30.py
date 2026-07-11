import math

from euler_lib import get_digits

# use the very lazy tactic of just making the loop maximum bigger and checking if the answer changed.
# then because 10000000 takes a very long time to run, just enter the number we have for that and see if its the right one
# turns out it is.
result = 0
for num in range(1000, 10000000):
    digits = get_digits(num)
    sum = 0
    for d in digits:
        sum += math.pow(d,5)
    if sum == num:
        result+=num
        print("num", num, "current sum", result)
print("result", result)
