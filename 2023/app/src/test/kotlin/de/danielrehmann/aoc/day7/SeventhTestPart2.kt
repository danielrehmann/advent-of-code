package de.danielrehmann.aoc.day7

import de.danielrehmann.aoc.Utils
import de.danielrehmann.aoc.day7.SeventhTestPart2.Card.*
import org.junit.jupiter.api.Test
import kotlin.test.assertEquals

class SeventhTestPart2 {
    val input = """
        32T3K 765
        T55J5 684
        KK677 28
        KTJJT 220
        QQQJA 483
    """.trimIndent()

    @Test
    fun canParseHand() {
        val given = "32T3K 765"

        val hand = parseLine(given)

        assertEquals(Hand(listOf(THREE, TWO, TEN, THREE, KING), 765), hand)
        assertEquals(Worth.PAIR, hand.worth)
    }

    @Test
    fun canOrderHands() {
        val hands = input.lines().map { parseLine(it) }

        assertEquals(
            listOf(
                Hand(cards = listOf(THREE, TWO, TEN, THREE, KING), bet = 765),
                Hand(cards = listOf(KING, KING, SIX, SEVEN, SEVEN), bet = 28),
                Hand(cards = listOf(TEN, FIVE, FIVE, JOKER, FIVE), bet = 684),
                Hand(cards = listOf(QUEEN, QUEEN, QUEEN, JOKER, ASS), bet = 483),
                Hand(cards = listOf(KING, TEN, JOKER, JOKER, TEN), bet = 220),
            ), hands.sorted()
        )

        val sorted = hands.sorted()
        assertEquals(5905, sorted.mapIndexed { index, hand -> (index + 1) * hand.bet }.sum())
    }

    @Test
    fun resultPart1() {
        val hands = Utils.readLinesFromFile("/7/input1.txt").map { parseLine(it) }
        val sorted = hands.sorted().associateWith { it.worth }
        assertEquals(253253225, sorted.keys.mapIndexed { index, hand -> (index + 1) * hand.bet }.sum())
        //wrong: 252677735
    }

    enum class Card {
        ASS,
        KING,
        QUEEN,
        TEN,
        NINE,
        EIGHT,
        SEVEN,
        SIX,
        FIVE,
        FOUR,
        THREE,
        TWO,
        JOKER;

        companion object {
            fun of(it: Char): Card {
                return when (it) {
                    'A' -> ASS
                    'K' -> KING
                    'Q' -> QUEEN
                    'J' -> JOKER
                    'T' -> TEN
                    '9' -> NINE
                    '8' -> EIGHT
                    '7' -> SEVEN
                    '6' -> SIX
                    '5' -> FIVE
                    '4' -> FOUR
                    '3' -> THREE
                    '2' -> TWO
                    else -> error("unkown card $it")
                }
            }
        }
    }

    enum class Worth {
        FIVE,
        FOUR,
        FULL_HOUSE,
        THREE,
        TWO_PAIRS,
        PAIR,
        HIGH_CARD
    }

    private fun parseLine(given: String): Hand {
        val (handString, betString) = given.split(" ").filter { it.isNotBlank() }
        return handString.map { Card.of(it) }.let { Hand(it, betString.toLong()) }
    }

    data class Hand(val cards: List<Card>, val bet: Long) : Comparable<Hand> {
        val worth: Worth = cards.groupingBy { it }.eachCount().let { heatMap ->
            val jokerAmount = heatMap.getOrDefault(JOKER, 0)
            val cardsWithoutJokers = heatMap.filterNot { it.key == JOKER }
            when {
                jokerAmount == 5 -> Worth.FIVE
                cardsWithoutJokers.values.any { it + jokerAmount >= 5 } -> Worth.FIVE
                cardsWithoutJokers.values.any { it + jokerAmount == 4 } -> Worth.FOUR
                isFullHouse(cardsWithoutJokers, jokerAmount) -> Worth.FULL_HOUSE
                cardsWithoutJokers.values.any { it + jokerAmount == 3 } -> Worth.THREE
                containsTwoPairs(cardsWithoutJokers.values, jokerAmount) -> Worth.TWO_PAIRS

                cardsWithoutJokers.values.any { it + jokerAmount == 2 } -> Worth.PAIR
                cardsWithoutJokers.values.all { it == 1 } -> Worth.HIGH_CARD
                else -> error("unknown card config $cards")
            }
        }

        private fun isFullHouse(
            cardsWithoutJokers: Map<Card, Int>,
            jokerAmount: Int
        ): Boolean {
            val triple = cardsWithoutJokers.filter { it.value == 3 }.count()
            val pairs = cardsWithoutJokers.filter { it.value == 2 }.count()
            if (triple > 0 && jokerAmount > 0) {
                error("should not be evaluated for full house already four of a kind")
            }
            if(pairs == 1 && jokerAmount > 1) {
                error("should not be evaluated for full house already four of a kind")
            }
            if (triple == 1 && pairs == 1) {
                return true
            }
            if (triple == 0 && pairs == 2 && jokerAmount == 1) {
                return true
            }
            return false

        }

        private fun containsTwoPairs(values: Collection<Int>, jokerAmount: Int) =
            values.fold(0) { acc, i ->
                if (i == 2) {
                    acc + 1
                } else {
                    acc
                }
            } == 2

        override fun compareTo(other: Hand): Int {
            return if (worth.ordinal < other.worth.ordinal) {
                1
            } else if (worth.ordinal > other.worth.ordinal) {
                -1
            } else {
                compareCards(other.cards)
            }
        }

        private fun compareCards(otherCards: List<Card>): Int {
            cards.forEachIndexed { index, card ->
                if (card.ordinal < otherCards[index].ordinal) {
                    return 1
                } else if (card.ordinal > otherCards[index].ordinal) {
                    return -1
                }
            }
            return 0
        }
    }

}