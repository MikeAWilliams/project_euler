import csv
import string


def get_data():
    words = []
    with open("0042_words.txt", mode="r", encoding="utf-8") as file:
        csv_reader = csv.reader(file)
        words = next(csv_reader)
    return words


def get_letter_score():
    letter_score = {}
    score = 1
    for letter in string.ascii_uppercase:
        letter_score[letter] = score
        score += 1
    return letter_score


def get_word_score(letter_score, word):
    total = 0
    for l in word:
        total += letter_score[l]
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
letter_score = get_letter_score()
triangles = get_triange_numbers(192)
print(triangles)

max = 0
for w in words:
    s = get_word_score(letter_score, w)
    if s > max:
        max = s
