package de.danielrehmann.aoc.day3

import de.danielrehmann.aoc.Utils
import org.junit.jupiter.api.Test
import kotlin.test.assertEquals

class ThirdOfDecember2ndTryTest {
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

    data class Num(val start: Int, val end: Int, val yPos: Int, val stringVal: String)
    data class Symbol(val pos: Int, val stringVal: String)

    @Test
    fun singleLineTest() {
        val line = "467..114..5"
        val numbers = getListOfNumbers(line)

        val symbols = line.mapIndexed { index, c ->
            if (!c.isDigit() && '.' != c) {
                Symbol(index, "$c")
            } else null
        }.filterNotNull()

        assertEquals(listOf(Num(0, 2, 0, "467"), Num(5, 7, 0, "114"), Num(10, 10, 0, "5")), numbers)
        assertEquals(listOf(), symbols)

        val singleSymbol = "...*......".mapIndexed { index, c ->
            if (!c.isDigit() && '.' != c) {
                Symbol(index, "$c")
            } else null
        }
        assertEquals(Symbol(3, "*"), singleSymbol[3])
    }

    private fun getListOfNumbers(
        line: String, yPos: Int = 0
    ): MutableList<Num> {
        val numbers = mutableListOf<Num>()
        var currentNumber: Num? = null
        line.forEachIndexed { index, point ->
            currentNumber = if (point.isDigit()) {
                currentNumber?.let {
                    Num(it.start, index, yPos, "${it.stringVal}$point")
                } ?: Num(index, index, yPos, "$point")
            } else {
                currentNumber?.let {
                    numbers.add(it)
                }
                null
            }
        }

        currentNumber?.let { numbers.add(it) }
        return numbers
    }

    data class Result(val neighbours: List<Num>, val nonNeighbours: List<Num>) {
        fun getSum(): Int = neighbours.sumOf { it.stringVal.toInt() }
    }

    @Test
    fun fullTest() {
        val result = getNeighbours(fullTestInput.lines())

        assertEquals(listOf("114", "58"), result.nonNeighbours.map { it.stringVal })

        assertEquals(4361, result.getSum())
    }


    @Test
    fun result() {
        val result = getNeighbours(Utils.readLinesFromFile("/3/input1.txt"))
        assertEquals(539590, result.getSum())
    }

    private fun getNeighbours(lines: List<String>): Result {
        val neighbours = mutableListOf<Num>()
        val nonNeighbours = mutableListOf<Num>()
        val numbers = lines.mapIndexed { index, it ->
            getListOfNumbers(it, index)
        }
        val symbols = lines.map {
            it.mapIndexed { index, c ->
                if (!c.isDigit() && '.' != c) {
                    Symbol(index, "$c")
                } else null
            }
        }

        numbers.forEachIndexed { yIndex, nums ->
            nums.forEach {
                if (isNeighbour(yIndex, it, symbols)) {
                    neighbours.add(it)
                } else {
                    nonNeighbours.add(it)
                }
            }
        }
        return Result(neighbours, nonNeighbours)
    }

    private fun isNeighbour(
        yIndex: Int, num: Num, symbols: List<List<Symbol?>>
    ): Boolean {
        return (num.start..num.end).map { hasSymbolNeighbour(yIndex, it, symbols) }.any { it }
    }

    private fun hasSymbolNeighbour(yIndex: Int, xIndex: Int, lists: List<List<Symbol?>>): Boolean {
        // yIndex - 1, xIndex-1
        // yIndex - 1, xIndex
        // yIndex - 1, xIndex+1
        // yIndex, xIndex-1
        // yIndex, xIndex+1
        // yIndex + 1, xIndex-1
        // yIndex + 1, xIndex
        // yIndex + 1, xIndex+1
        val symbolNeighbour = listOf(lists.getOrNull(yIndex - 1)?.getOrNull(xIndex - 1)?.let { true } ?: false,
            lists.getOrNull(yIndex - 1)?.getOrNull(xIndex)?.let { true } ?: false,
            lists.getOrNull(yIndex - 1)?.getOrNull(xIndex + 1)?.let { true } ?: false,
            lists.getOrNull(yIndex)?.getOrNull(xIndex - 1)?.let { true } ?: false,
            lists.getOrNull(yIndex)?.getOrNull(xIndex + 1)?.let { true } ?: false,
            lists.getOrNull(yIndex + 1)?.getOrNull(xIndex - 1)?.let { true } ?: false,
            lists.getOrNull(yIndex + 1)?.getOrNull(xIndex)?.let { true } ?: false,
            lists.getOrNull(yIndex + 1)?.getOrNull(xIndex + 1)?.let { true } ?: false)
        return symbolNeighbour.any { it }
    }
    // PART 2

    data class Point(val y: Int, val x: Int)

    @Test
    fun calculateGearRatio() {
        val result = getNeighbours(fullTestInput.lines())

        val mapOfNeighbours = result.neighbours.flatMap { num ->
            (num.start..num.end).map {
                Point(
                    num.yPos, it
                ) to num
            }
        }.associate { it }

        assertEquals(467, mapOfNeighbours.get(Point(0, 0))?.stringVal?.toInt())
        assertEquals(467, mapOfNeighbours.get(Point(0, 1))?.stringVal?.toInt())
        assertEquals(467, mapOfNeighbours.get(Point(0, 2))?.stringVal?.toInt())
        assertEquals(null, mapOfNeighbours.get(Point(1, 0))?.stringVal?.toInt())

        val potentialGears = fullTestInput.lines().flatMapIndexed { yIndex, line ->
            line.mapIndexed { index, c ->
                if (c == '*') {
                    Point(yIndex, index)
                } else null
            }.filterNotNull()
        }

        assertEquals(listOf(Point(y = 1, x = 3), Point(y = 4, x = 3), Point(y = 8, x = 5)), potentialGears)

        val point = Point(y = 1, x = 3)
        val potentialNeighbours = (-1..1).flatMap { x ->
            (-1..1).map { y ->
                Point(point.y + y, point.x + x)
            }
        }
        assertEquals(
            listOf(
                Point(y = 0, x = 2),
                Point(y = 1, x = 2),
                Point(y = 2, x = 2),
                Point(y = 0, x = 3),
                Point(y = 1, x = 3),
                Point(y = 2, x = 3),
                Point(y = 0, x = 4),
                Point(y = 1, x = 4),
                Point(y = 2, x = 4)
            ), potentialNeighbours
        )

        val sum = getGearSum(potentialGears, mapOfNeighbours)
        assertEquals(467835, sum)
    }

    @Test
    fun part2Result() {
        val result = getNeighbours(Utils.readLinesFromFile("/3/input1.txt"))
        val mapOfNeighbours = result.neighbours.flatMap { num ->
            (num.start..num.end).map {
                Point(
                    num.yPos, it
                ) to num
            }
        }.associate { it }
        val potentialGears = Utils.readLinesFromFile("/3/input1.txt").flatMapIndexed { yIndex, line ->
            line.mapIndexed { index, c ->
                if (c == '*') {
                    Point(yIndex, index)
                } else null
            }.filterNotNull()
        }
        val sum = getGearSum(potentialGears, mapOfNeighbours)
        assertEquals(80703636, sum)
    }

    private fun getGearSum(
        potentialGears: List<Point>,
        mapOfNeighbours: Map<Point, Num>
    ) = potentialGears.sumOf {
        val allNeighbours = (-1..1).flatMap { x ->
            (-1..1).map { y -> Point(it.y + y, it.x + x) }
        }.mapNotNull { mapOfNeighbours[it] }.distinct()

        println(allNeighbours)

        if (allNeighbours.size == 2) {
            allNeighbours[0].stringVal.toInt() * allNeighbours[1].stringVal.toInt()
        } else {
            0
        }
    }


}