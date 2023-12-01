package de.danielrehmann.aoc

import de.danielrehmann.aoc.Utils.Companion.readLinesFromFile
import org.junit.jupiter.api.Test
import kotlin.test.assertEquals

class FirstOfDecemberSecondTest {
    val testInput = """
        two1nine
        eightwothree
        abcone2threexyz
        xtwone3four
        4nineeightseven2
        zoneight234
        7pqrstsixteen
    """.trimIndent()
    val expectedOutput = listOf(29, 83, 13, 24, 42, 14, 76)
    val expectedSum = 281

    @Test
    fun parseSingleLine() {
        val given = "two1nine"

        assertEquals(29, given.getFirstAndLastDigitWithTextual())
    }

    @Test
    fun allLinesMatch() {
        assertEquals(expectedOutput, testInput.lines().map { it.getFirstAndLastDigitWithTextual() })
        assertEquals(expectedSum, testInput.lines().sumOf { it.getFirstAndLastDigitWithTextual() })
    }

    @Test
    fun firstAnswer() {
        val answer = readLinesFromFile("/1/input1.txt").sumOf { it.getFirstAndLastDigitWithTextual() }
        assertEquals(55413, answer)
    }
}

val numbers = listOf(
    "one",
    "two",
    "three",
    "four",
    "five",
    "six",
    "seven",
    "eight",
    "nine",
    "1",
    "2",
    "3",
    "4",
    "5",
    "6",
    "7",
    "8",
    "9"
)

fun String.getFirstAndLastDigitWithTextual(): Int {
    val firstAndLastNumbers =
        (findAnyOf(numbers)?.second ?: error("No First Number")) to (findLastAnyOf(numbers)?.second
            ?: error("No Last Number"))
    firstAndLastNumbers.toConcatenatedInt()
    return "${findAnyOf(numbers)?.second?.textDigitToDigit()}${findLastAnyOf(numbers)?.second?.textDigitToDigit()}".toInt()
}

private fun Pair<String, String>.toConcatenatedInt(): Int {
    val stringConcat = first.textDigitToDigit() + second.textDigitToDigit()
    return stringConcat.toInt()

}

private fun String.textDigitToDigit(): String = when (this) {
    "one", "1" -> "1"
    "two", "2" -> "2"
    "three", "3" -> "3"
    "four", "4" -> "4"
    "five", "5" -> "5"
    "six", "6" -> "6"
    "seven", "7" -> "7"
    "eight", "8" -> "8"
    "nine", "9" -> "9"
    else -> error("$this ist not parsable to number")
}
