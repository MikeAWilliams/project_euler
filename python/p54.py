from dataclasses import dataclass
from enum import Enum

demo_hands = [
    "5H 5C 6S 7S KD 2C 3S 8S 8D TD",
    "5D 8C 9S JS AC 2C 5C 7D 8S QH",
    "2D 9C AS AH AC 3D 6D 7D TD QD",
    "4D 6S 9H QH QC 3D 6D 7H QD QS",
    "2H 2D 4C 4D 4S 3C 3D 3S 9S 9D",
]


class Suit(Enum):
    DIAMONDS = 1
    HEARTS = 2
    SPADES = 3
    CLUBS = 4


class Value(Enum):
    TWO = 2
    THREE = 3
    FOUR = 4
    FIVE = 5
    SIX = 6
    SEVEN = 7
    EIGHT = 8
    NINE = 9
    TEN = 10
    JACK = 11
    QUEEN = 12
    KING = 13
    ACE = 14


class Card:
    def __init__(self, letters):
        if len(letters) != 2:
            raise Exception("only 2 is valid for lenthg of letters")
        match letters[0]:
            case "2":
                self.val = Value.TWO
            case "3":
                self.val = Value.THREE
            case "4":
                self.val = Value.FOUR
            case "5":
                self.val = Value.FIVE
            case "6":
                self.val = Value.SIX
            case "7":
                self.val = Value.SEVEN
            case "8":
                self.val = Value.EIGHT
            case "9":
                self.val = Value.NINE
            case "T":
                self.val = Value.TEN
            case "J":
                self.val = Value.JACK
            case "Q":
                self.val = Value.QUEEN
            case "K":
                self.val = Value.KING
            case "A":
                self.val = Value.ACE
            case _:
                raise Exception(f"Invalid value {letters[0]}")

        match letters[1]:
            case "D":
                self.suit = Suit.DIAMONDS
            case "H":
                self.suit = Suit.HEARTS
            case "S":
                self.suit = Suit.SPADES
            case "C":
                self.suit = Suit.CLUBS
            case _:
                raise Exception(f"Invalid suit {letters[1]}")


def string_to_hands(string):
    cards = string.split()
    if len(cards) != 10:
        raise Exception("It should always be 10")
    index = 0
    p1 = []
    p2 = []
    for c in cards:
        if index < 5:
            p1.append(Card(c))
        else:
            p2.append(Card(c))
        index += 1
    return p1, p2


print(string_to_hands(demo_hands[0]))
