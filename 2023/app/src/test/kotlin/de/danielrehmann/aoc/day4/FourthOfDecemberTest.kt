package de.danielrehmann.aoc.day4

import de.danielrehmann.aoc.Utils
import org.junit.jupiter.api.Test
import kotlin.test.assertEquals

class FourthOfDecemberTest {
    val testInput = """
        Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
        Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
        Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
        Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
        Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
        Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
    """.trimIndent()

    @Test
    fun calculateSingleCard() {
        val given = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"

        val actual = Card.of(given)
        assertEquals(Card(1, setOf(41, 48, 83, 86, 17), listOf(83, 86, 6, 31, 17, 9, 48, 53)), actual)

        assertEquals(8, actual.points)
    }

    @Test
    fun calculateAllCards() {
        assertEquals(13, testInput.lines().map { Card.of(it) }.sumOf { it.points })
    }

    @Test
    fun result() {
        assertEquals(17803, Utils.readLinesFromFile("/4/input1.txt").map { Card.of(it) }.sumOf { it.points })
    }

    // 2nd part

    @Test
    fun singleCardNextWinningCards() {
        val given = "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"

        val actual = Card.of(given)

        assertEquals(listOf(2, 3, 4, 5), actual.wonScratchCards)
    }

    @Test
    fun calculateAllCardsWinningCards() {
        val actual = testInput.lines().map { Card.of(it) }
        val cards = actual.fold(emptyList<Int>()) {
                acc, card ->
            val processings = acc.count { it == card.id } + 1
            acc + card.id + (1..processings)
                .flatMap { card.wonScratchCards }
        }.count()



        assertEquals(30, cards)
    }

    @Test
    fun resultWinningCards() {
        val actual = Utils.readLinesFromFile("/4/input1.txt").map { Card.of(it) }

        val cards = actual.fold(emptyList<Int>()) {
                acc, card ->
            val processings = acc.count { it == card.id } + 1
            acc + card.id + (1..processings)
                .flatMap { card.wonScratchCards }
        }.count()



        assertEquals(5554894, cards)
    }
}

data class Card(val id: Int, val winningNumbers: Set<Int>, val numbers: List<Int>) {
    val points: Int
        get() = (1..matches).fold(0) { acc, _ -> if (acc == 0) 1 else acc * 2 }
    private val matches = numbers.count { winningNumbers.contains(it) }
    val wonScratchCards: List<Int> = (1..matches).map { id + it }

    companion object {
        private val cardRegex = Regex("^Card\\s+(\\d+): ([\\d\\s]+) \\| ([\\d\\s]+)\$")

        fun of(text: String): Card = cardRegex.matchEntire(text)?.let {
            val id = it.groupValues[1].toInt()
            val winningNumbers = it.groupValues[2].split(" ").filterNot { it.isEmpty() }.map { it.toInt() }
            val numbers = it.groupValues[3].split(" ").filterNot { it.isEmpty() }.map { it.toInt() }
            Card(id, winningNumbers.toSet(), numbers)
        } ?: error("Could not parse $text")

    }
}