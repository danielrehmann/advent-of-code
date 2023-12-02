package de.danielrehmann.aoc.day1

import de.danielrehmann.aoc.Utils.Companion.readLinesFromFile
import org.junit.jupiter.api.Test
import kotlin.test.assertEquals

class FirstOfDecemberTest {
    val testInput = """
        1abc2
        pqr3stu8vwx
        a1b2c3d4e5f
        treb7uchet
    """.trimIndent()
    val expectedOutput = listOf(12, 38, 15, 77)
    val expectedSum = 142

    @Test
    fun parseFirstAndLastDigitFromLine() {
        // given
        val input = "1abc2"
        // when
        val result = input.getFirstAndLastDigit()

        // then
        assertEquals(expected = 12, actual = result)
    }

    @Test
    fun allLinesMatch() {
        assertEquals(expectedOutput, testInput.lines().map { it.getFirstAndLastDigit() })
        assertEquals(expectedSum, testInput.lines().sumOf { it.getFirstAndLastDigit() })
    }

    @Test
    fun firstAnswer() {
        val answer = readLinesFromFile("/1/input1.txt").sumOf { it.getFirstAndLastDigit() }
        assertEquals(55712, answer)
    }
}
fun String.getFirstAndLastDigit() = "${first { it.isDigit() }}${reversed().first { it.isDigit() }}".toInt()
