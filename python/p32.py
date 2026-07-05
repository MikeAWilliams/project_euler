import math


def get_digits(n):
    result = []
    while n > 0:
        digit = n % 10
        result.append(digit)
        n = n // 10
    return result

def has_zeros(digit_list):
    for a in digit_list:
        if a == 0:
            return True
    return False

def has_duplicates(digit_list):
    unique_set = set()
    for d in digit_list:
        unique_set.add(d)
    return len(unique_set) != len(digit_list)

def get_num_digits(num):
    #return int(math.log(num,10))+1
    return len(str(num))

def is_pandigital(num, required_digits):
    real_digits = get_num_digits(num)
    if real_digits != required_digits:
        return False

    digits = get_digits(num)
    if len(digits) != real_digits:
        raise Exception(f"digits didn't work, num {num}, digits {digits}")

    if has_zeros(digits):
        return False
    if has_duplicates(digits):
        return False
    return True

def pad_number_to_digits(num, required_digits):
    result = num
    for i in range(required_digits):
        result = result * 10
    return result

def product_and_args_are_pandigital9(a, b, product):
    if a <= 0 or b <= 0 or product <= 0:
        return False
    padded_product = pad_number_to_digits(product, get_num_digits(a) + get_num_digits(b))
    padded_a = pad_number_to_digits(a, get_num_digits(b))

    smoosh = padded_product + padded_a + b
    return(is_pandigital(smoosh, 9))

def explore_digit_space():
    all_a = set()
    all_b = set()
    for a in range(1, 100000000):
        if a % 100 == 0:
            print(a)
        for b in range(1, 100000000):
            p = a * b
            digits_a = get_num_digits(a)
            digits_b = get_num_digits(b)
            digits_p = get_num_digits(p)
            sum = digits_a + digits_b + digits_p
            if sum == 9:
                all_a.add(a)
                all_b.add(b)
                #print(f"a:{a}, b:{b}, p:{p}, da:{digits_a}, db:{digits_b}, dp:{digits_p}")
            if sum > 9:
                break
    print("min a", min(all_a), "max a", max(all_a))
    print("min b", min(all_b), "max b", max(all_b))
#print("explore the digits space")
#explore_digit_space()
result = 0
products = set()

for a in range(1,10000):
   for b in range(1,10000):
       p = a*b
       if product_and_args_are_pandigital9(a,b,p):
           if p not in products:
               result += p
               products.add(p)
       if get_num_digits(p) + get_num_digits(a) + get_num_digits(b) > 9:
            break
print(result)
# was 34# to square after entering this solution
