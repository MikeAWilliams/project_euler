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

strait_not = ["2D 3H 4S 5C 6D 2H 4D 6S 8D 9C"]

mike_hands = [
    "TD JD QD KD AD 2H 3H 4H 5H 6H",
    "2D JD QD KD AD 2H 2D 2S 2C 6H",
    "2D 2D 3S 3S 3S 2H 3D 4S 5C 6H",
    "2D 2D 2S 3S 4S 2H 2D 4S 4C 6H",
    "2D 2D 3S 5S 4S 2H 4D 5S 6C 7H",
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


class HandRank(IntEnum):
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


def is_flush(hand):
    suits = set()
    for c in hand:
        suits.add(c.suit)
    return len(suits) == 1


def is_straight(hand):
    # assums hand was sorted
    vlast = hand[0].val
    for i in range(1, 5):
        if hand[i].val != vlast + 1:
            return False
        vlast = hand[i].val
    return True


def count_values(hand):
    count = {}
    for c in hand:
        if c.val in count:
            count[c.val] += 1
        else:
            count[c.val] = 1
    vals = list(count.values())
    vals.sort()
    return vals


def rank_hand(hand):
    hand.sort(key=lambda c: c.val)
    hand_is_flush = is_flush(hand)
    hand_is_straight = is_straight(hand)
    if hand_is_flush:
        if hand[0].val == Value.TEN:
            return HandRank.ROYAL_FLUSH
        if hand_is_straight:
            return HandRank.STRAIGHT_FLUSH
        return HandRank.FLUSH
    count_tup = count_values(hand)
    if len(count_tup) == 2:
        if count_tup[0] == 1:
            return HandRank.FOUR_OF_A_KIND
        return HandRank.FULL_HOUSE
    if hand_is_straight:
        return HandRank.STRAIGHT
    if len(count_tup) == 3:
        if count_tup[2] == 3:
            return HandRank.THREE_OF_A_KIND
        return HandRank.TWO_PAIRS
    if len(count_tup) == 4:
        return HandRank.ONE_PAIR
    return HandRank.HIGH_CARD

def sort_values_for_rank(rank, hand):
    # assum hand is sorted
    vals = [c.val for c in hand]
    match(rank):
        case HandRank.HIGH_CARD:
           return vals[::-1] # return values in highest to lowest
        case HandRank.ONE_PAIR:
        case HandRank.TWO_PAIRS:
        case HandRank.THREE_OF_A_KIND:
        case HandRank.STRAIGHT:
           return vals[::-1] # return values in highest to lowest
        case HandRank.FLUSH:
           return vals[::-1] # return values in highest to lowest
        case HandRank.FULL_HOUSE:
        case HandRank.FOUR_OF_A_KIND:
        case HandRank.STRAIGHT_FLUSH:
           return vals[::-1] # return values in highest to lowest
        case HandRank.ROYAL_FLUSH:
           return vals[::-1] # return values in highest to lowest

working_hands = demo_hands
# working_hands = mike_hands
# working_hands = strait_not
# working_hands = promblem_hands[0:5]

p1_wins = 0
for l in working_hands:
    print(l)
    h1, h2 = string_to_hands(l)
    h1_rank = rank_hand(h1)
    h2_rank = rank_hand(h2)
    if h1_rank > h2_rank:
        p1_wins += 1
    elif h1_rank == h2_rank:
        # determine the card in the ranking that counts
        hand_by_rank_1 = sort_values_for_rank(h1_rank, h1)
        hand_by_rank_2 = sort_values_for_rank(h2_rank, h2)
        for index in range(5):
            if hand_by_rank_1[index] > hand_by_rank_2[index]:
                p1_wins += 1
                break
            elif hand_by_rank_2[index] > hand_by_rank_1[index]:
                # p2 wins
                break
            # else it must be a tie keep looking
print(p1_wins)
