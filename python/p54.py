from enum import Enum, IntEnum


def read_hands():
    with open("0054_poker.txt") as f:
        return [line.strip() for line in f]


promblem_hands = read_hands()

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

    @property
    def symbol(self):
        return {
            Suit.DIAMONDS: "D",
            Suit.HEARTS: "H",
            Suit.SPADES: "S",
            Suit.CLUBS: "C",
        }[self]


class Value(IntEnum):
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

    @property
    def symbol(self):
        return {10: "T", 11: "J", 12: "Q", 13: "K", 14: "A"}.get(
            self.value, str(self.value)
        )


class HandRank(Enum):
    HIGH_CARD = 1
    ONE_PAIR = 2
    TWO_PAIRS = 3
    THREE_OF_A_KIND = 4
    STRAIGHT = 5
    FLUSH = 6
    FULL_HOUSE = 7
    FOUR_OF_A_KIND = 8
    STRAIGHT_FLUSH = 9
    ROYAL_FLUSH = 10


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

    def __repr__(self):
        return f"{self.val.symbol}{self.suit.symbol}"


def string_to_hands(string):
    cards = string.split()
    if len(cards) != 10:
        raise Exception("It should always be 10")
    p1 = []
    p2 = []
    for index, c in enumerate(cards):
        if index < 5:
            p1.append(Card(c))
        else:
            p2.append(Card(c))
    return p1, p2


def rank_hand(hand):
    hand.sort(key=lambda c: c.val)
    return HandRank.FLUSH


working_hands = demo_hands

for l in working_hands:
    h1, h2 = string_to_hands(l)
    h1_rank = rank_hand(h1)
    h2_rank = rank_hand(h2)
    print(h1, h1_rank, h2, h2_rank)
