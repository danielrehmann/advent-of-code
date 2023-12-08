package de.danielrehmann.aoc.day7

import de.danielrehmann.aoc.Utils
import de.danielrehmann.aoc.day7.SeventhTest.Card.*
import org.junit.jupiter.api.Test
import kotlin.test.assertEquals

class SeventhTest {
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
                Hand(cards = listOf(TEN, FIVE, FIVE, JESTER, FIVE), bet = 684),
                Hand(cards = listOf(KING, KING, SIX, SEVEN, SEVEN), bet = 28),
                Hand(cards = listOf(KING, TEN, JESTER, JESTER, TEN), bet = 220),
                Hand(cards = listOf(QUEEN, QUEEN, QUEEN, JESTER, ASS), bet = 483)
            ), hands
        )

        assertEquals(
            listOf(
                Hand(cards = listOf(QUEEN, QUEEN, QUEEN, JESTER, ASS), bet = 483),
                Hand(cards = listOf(TEN, FIVE, FIVE, JESTER, FIVE), bet = 684),
                Hand(cards = listOf(KING, KING, SIX, SEVEN, SEVEN), bet = 28),
                Hand(cards = listOf(KING, TEN, JESTER, JESTER, TEN), bet = 220),
                Hand(cards = listOf(THREE, TWO, TEN, THREE, KING), bet = 765),
            ), hands.sorted().reversed()
        )

        assertEquals(6440, hands.sorted().mapIndexed { index, hand -> (index + 1) * hand.bet }.sum())
    }

    @Test
    fun resultPart1() {
        val hands = Utils.readLinesFromFile("/7/input1.txt").map { parseLine(it) }
        assertEquals(253638586, hands.sorted().mapIndexed { index, hand -> (index + 1) * hand.bet }.sum())
    }

    enum class Card {
        ASS,
        KING,
        QUEEN,
        JESTER,
        TEN,
        NINE,
        EIGHT,
        SEVEN,
        SIX,
        FIVE,
        FOUR,
        THREE,
        TWO;

        companion object {
            fun of(it: Char): Card {
                return when (it) {
                    'A' -> ASS
                    'K' -> KING
                    'Q' -> QUEEN
                    'J' -> JESTER
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
            when {
                heatMap.values.any { it == 5 } -> Worth.FIVE
                heatMap.values.any { it == 4 } -> Worth.FOUR
                heatMap.values.any { it == 3 } && heatMap.values.any { it == 2 } -> Worth.FULL_HOUSE
                heatMap.values.any { it == 3 } -> Worth.THREE
                containsTwoPairs(heatMap.values) -> Worth.TWO_PAIRS

                heatMap.values.any { it == 2 } -> Worth.PAIR
                heatMap.values.all { it == 1 } -> Worth.HIGH_CARD
                else -> error("unknown card config $cards")
            }
        }

        private fun containsTwoPairs(values: Collection<Int>) =
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