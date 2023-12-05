package de.danielrehmann.aoc.day5

import de.danielrehmann.aoc.Utils
import org.junit.jupiter.api.Test
import kotlin.test.assertEquals

class Fifth2nTryTest {

    val input = """
        seeds: 79 14 55 13

        seed-to-soil map:
        50 98 2
        52 50 48

        soil-to-fertilizer map:
        0 15 37
        37 52 2
        39 0 15

        fertilizer-to-water map:
        49 53 8
        0 11 42
        42 0 7
        57 7 4

        water-to-light map:
        88 18 7
        18 25 70

        light-to-temperature map:
        45 77 23
        81 45 19
        68 64 13

        temperature-to-humidity map:
        0 69 1
        1 0 69

        humidity-to-location map:
        60 56 37
        56 93 4
    """.trimIndent()

    @Test
    fun parseContent() {
        val (seeds, steps) = Parser.parse(input.lines())
        val step = steps.first()

        val seedRanges = seeds.map { it.rangeTo(it) }

        assertEquals(listOf(52L..52), step.map(50L..50))
        assertEquals(listOf(82L..82), steps.fold(listOf(79L..79)) { acc, s -> s.map(acc) })
        assertEquals(35, steps.fold(seedRanges) { acc, s -> s.map(acc) }.minOf { it.first })
    }

    @Test
    fun parseContentResult() {
        val (seeds, steps) = Parser.parse(Utils.readLinesFromFile("/5/input1.txt"))

        val seedRanges = seeds.map { it.rangeTo(it) }

        assertEquals(57075758, steps.fold(seedRanges) { acc, s -> s.map(acc) }.minOf { it.first })
    }

    @Test
    fun rangeTest() {
        val (seeds, steps) = Parser.parse(input.lines())

        val seedRanges = seeds.chunked(2).map { (first, second) -> first..<first + second }

        //assertEquals(listOf(52L..52), step.map(50L..50))
        assertEquals(listOf(82L..82), steps.fold(listOf(79L..79)) { acc, s -> s.map(acc) })
        assertEquals(46, steps.fold(seedRanges) { acc, s -> s.map(acc) }.minOf { it.first })
    }

    @Test
    fun rangeResultTest() {
        val (seeds, steps) = Parser.parse(Utils.readLinesFromFile("/5/input1.txt"))

        val seedRanges = seeds.chunked(2).map { (first, second) -> first..<first + second }

        assertEquals(31161857, steps.fold(seedRanges) { acc, s -> s.map(acc) }.minOf { it.first })
    }

}

data class Step(val name: String, val mappers: List<Mapper>) {
    fun map(longRange: LongRange): List<LongRange> {
        val matchList = mappers.mapNotNull { it.getMatch(longRange) }
        val foundValues = matchList.sortedBy { it.found.first }.flatMap { listOf(it.found.first, it.found.last) }
        val changedValues = matchList.map { it.mapped }

        val unchangedRanges = if (foundValues.isNotEmpty()) {
            foundValues.addFirst(longRange.first)
            foundValues.addLast(longRange.last)
            val res = foundValues.chunked(2).map { (first, last) -> first..<last-1 }
                .filterNot { it.isEmpty() }
            res
        } else {
            listOf(longRange)
        }

        return changedValues + unchangedRanges
    }

    fun map(longRanges: List<LongRange>): List<LongRange> {
        return longRanges.flatMap { map(it) }
    }

    companion object {
        fun of(strings: List<String>): Step {
            return Step(
                strings.first(),
                strings.subList(1, strings.size).let { stringList ->
                    stringList.map { s ->
                        val (destinationRange, sourceRange, rangeLength) = s.split(" ").map { it.toLong() }
                        Mapper.of(destinationRange, sourceRange, rangeLength)
                    }
                })
        }
    }
}

data class Mapper(val begin: Long, val end: Long, val shift: Long) {
    fun getMatch(rangeLength: LongRange): Match? = if (rangeLength.first <= end && rangeLength.last >= begin) {
        val found = maxOf(rangeLength.first, begin)..minOf(rangeLength.last, end)
        Match(found, found.first + shift..found.last + shift)
    } else {
        null
    }

    companion object {
        fun of(destinationRange: Long, sourceRange: Long, rangeLength: Long): Mapper {
            val begin = sourceRange
            val end = sourceRange + rangeLength - 1
            val shift = -(sourceRange - destinationRange)
            return Mapper(begin, end, shift)
        }
    }

}

data class Match(val found: LongRange, val mapped: LongRange)

data class Parser(val parts: List<List<String>> = emptyList(), val lastWasLinebreak: Boolean = true) {
    fun add(string: String): Parser {
        return if (string.isBlank()) {
            Parser(parts, true)
        } else {
            parts.toMutableList().let {
                if (lastWasLinebreak) {
                    it.add(listOf(string))
                } else {
                    it[it.lastIndex] = it.last().plus(string)
                }
                Parser(it, false)
            }
        }
    }

    companion object {
        private val seedRegex = Regex("^seeds: ([\\d|\\s]+)\$")
        fun parse(lines: List<String>): Pair<List<Long>, List<Step>> {
            val fold = lines.fold(Parser()) { acc, s -> acc.add(s) }
            return fold.parts.first().first()
                .let {
                    seedRegex.matchEntire(it)!!.groupValues[1].split(" ").map { s -> s.toLong() }
                } to fold.parts.subList(1, fold.parts.size).map {
                Step.of(it)
            }
        }
    }
}