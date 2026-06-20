# unsure of repeated numbers
def string_are_anagrams(str1, str2):
    if len(str1) != len(str2):
        return False
    for c1 in str1:
        found_in_2 = False
        for c2 in str2:
            if c2 == c1:
                found_in_2 = True
                break
        if not found_in_2:
            return False
    return True


print(string_are_anagrams("11", "22"))
print(string_are_anagrams("12", "13"))
print(string_are_anagrams("9", "18"))
print(string_are_anagrams("12", "21"))

n = 125874
n2 = n * 2
n_str = str(n)
n2_str = str(n2)
print(string_are_anagrams(n_str, n2_str))


for x in range(1, 2000000):
    x_str = str(x)
    x2_str = str(x * 2)
    if not string_are_anagrams(x_str, x2_str):
        continue
    x3_str = str(x * 3)
    if not string_are_anagrams(x_str, x3_str):
        continue
    x4_str = str(x * 4)
    if not string_are_anagrams(x_str, x4_str):
        continue
    x5_str = str(x * 5)
    if not string_are_anagrams(x_str, x5_str):
        x5_str = str(x * 5)
    x6_str = str(x * 6)
    if not string_are_anagrams(x_str, x6_str):
        continue
    print(x)
    break
print("done")
