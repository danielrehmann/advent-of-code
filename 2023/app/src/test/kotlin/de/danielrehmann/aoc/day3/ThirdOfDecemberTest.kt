package de.danielrehmann.aoc.day3

import de.danielrehmann.aoc.day3.ThirdOfDecemberTest.Point.Empty
import org.junit.jupiter.api.Test
import kotlin.test.assertEquals

class ThirdOfDecemberTest {

    sealed interface Point {

        data object Empty : Point

        data class Symbol(val value: Char) : Point

        data class Number(val value: Int) : Point
    }

    val fullTestInput = """
        467..114..
        ...*......
        ..35..633.
        ......#...
        617*......
        .....+.58.
        ..592.....
        ......755.
        ...$.*....
        .664.598..
    """.trimIndent()

    @Test
    fun canConvertLineToListOfObjects() {
        val given = "467..114.."
        val list = mapLine(given)
        assertEquals(
            listOf(
                '4', '6', '7', null, null, '1', '1', '4', null, null
            ).map { c ->
                c?.let { Point.Number(it.digitToInt()) } ?: Empty
            }, list
        )
    }

    @Test
    fun canConvertLineWithSymbolToListOfObjects() {
        val given = "...*......"
        val list = mapLine(given)
        assertEquals(
            listOf(
                Empty, Empty, Empty, Point.Symbol('*'), Empty, Empty, Empty, Empty, Empty, Empty
            ), list
        )
    }

    @Test
    fun canMapFullInputToListOfList(){
        val lists = getTestInputLists()

        assertEquals(Point.Number(4), lists[0][0])
        assertEquals(Empty, lists[1][0])
        assertEquals(Point.Symbol('*'), lists[1][3])
    }

    @Test
    fun findAllSymbolNeighboringValues() {
        val lists = getTestInputLists()
        val neighbours = mutableListOf<Point.Number>()
        val nonNeighbours = mutableListOf<Point.Number>()

        lists.forEachIndexed{ yIndex, points ->
            points.forEachIndexed { xIndex, point ->
                if (point is Point.Number) {
                    if (hasSymbolNeighbour(yIndex, xIndex, lists)) {
                        neighbours.add(point)
                    } else {
                        nonNeighbours.add(point)
                    }
                }
            }
        }

        //assertEquals(listOf(Point.Number(114), Point.Number(58)), nonNeighbours)
    }

    private fun hasSymbolNeighbour(yIndex: Int, xIndex: Int, lists: List<List<Point>>): Boolean {
        // yIndex - 1, xIndex-1
        // yIndex - 1, xIndex
        // yIndex - 1, xIndex+1
        // yIndex, xIndex-1
        // yIndex, xIndex+1
        // yIndex + 1, xIndex-1
        // yIndex + 1, xIndex
        // yIndex + 1, xIndex+1
        val symbolNeighbour = listOf(
        lists.getOrNull(yIndex - 1)?.getOrNull(xIndex - 1)?.let { it is Point.Symbol } ?: false,
        lists.getOrNull(yIndex - 1)?.getOrNull(xIndex)?.let { it is Point.Symbol } ?: false,
        lists.getOrNull(yIndex - 1)?.getOrNull(xIndex + 1)?.let { it is Point.Symbol } ?: false,
        lists.getOrNull(yIndex)?.getOrNull(xIndex - 1)?.let { it is Point.Symbol } ?: false,
        lists.getOrNull(yIndex)?.getOrNull(xIndex + 1)?.let { it is Point.Symbol } ?: false,
        lists.getOrNull(yIndex + 1)?.getOrNull(xIndex - 1)?.let { it is Point.Symbol } ?: false,
        lists.getOrNull(yIndex + 1)?.getOrNull(xIndex)?.let { it is Point.Symbol } ?: false,
        lists.getOrNull(yIndex + 1)?.getOrNull(xIndex + 1)?.let { it is Point.Symbol } ?: false)

        return symbolNeighbour.any { it }
    }

    private fun getTestInputLists() = fullTestInput.lines().map { mapLine(it) }

    private fun mapLine(given: String) = given.map { c ->
        when {
            c == '.' -> {
                Empty
            }

            c.isDigit() -> {
                Point.Number(c.digitToInt())
            }

            else -> {
                Point.Symbol(c)
            }
        }
    }

}