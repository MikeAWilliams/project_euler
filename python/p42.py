import csv
import string
from bisect import bisect_left


def binary_search(sorted_list, target):
    index = bisect_left(sorted_list, target)

    if index < len(sorted_list) and sorted_list[index] == target:
        return index
    return -1


def get_data():
    words = []
    with open("0042_words.txt", mode="r", encoding="utf-8") as file:
        csv_reader = csv.reader(file)
        words = next(csv_reader)
    return words


AVALUE = ord("A") - 1


def get_word_score(word):
    total = 0
    for l in word:
        total += ord(l) - AVALUE
    return total


def get_triangle(n):
    return int(0.5 * n * (n + 1))


def get_triange_numbers(max):
    result = []
    i = 1
    t = get_triangle(i)
    while t <= max:
        result.append(t)
        i += 1
        t = get_triangle(i)

    return result


words = get_data()
triangles = get_triange_numbers(192)


count = 0
for w in words:
    s = get_word_score(w)
    index = binary_search(triangles, s)
    if index >= 0:
        count += 1
print(count)
