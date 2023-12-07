package de.danielrehmann.aoc.day6

import org.junit.jupiter.api.Test
import java.util.concurrent.atomic.AtomicLong
import java.util.regex.Pattern
import kotlin.test.assertEquals

class SixthTestTest {

    val input = """
        Time:      7  15   30
        Distance:  9  40  200
    """.trimIndent()

    val realInput = """
        Time:        60     94     78     82
        Distance:   475   2138   1015   1650
    """.trimIndent()

    @Test
    fun parseInput() {
        val inpu = parse(input.lines())
        assertEquals(7L to 9L, inpu[0])
        assertEquals(15L to 40L, inpu[1])
        assertEquals(30L to 200L, inpu[2])
    }

    @Test
    fun findFasterRuns() {
        val run = 7L to 9L

        val res = calculateFaster(run.first, run.second)

        assertEquals(listOf(2L, 3, 4, 5), res)
    }

    @Test
    fun findInputTest() {
        val run = parse(input.lines())

        val res = run.map { calculateFaster(it.first, it.second).count() }.reduce { acc, i -> acc * i }

        assertEquals(288, res)
    }

    @Test
    fun findResultTest() {
        val run = parse(realInput.lines())

        val res = run.map { calculateFaster(it.first, it.second).count() }.reduce { acc, i -> acc * i }

        assertEquals(345015, res)
    }

    // part2
    @Test
    fun parseSingleRace() {
        val res = parseSingleRace(input.lines())
        assertEquals(71530L to 940200L, res)
    }

    @Test
    fun inputTestSingleRace() {
        val res = parseSingleRace(input.lines()).let {
            calculateFaster(it.first, it.second).count()
        }
        assertEquals(71503, res)
    }

    @Test
    fun resultTestSingleRace() {
        val res = parseSingleRace(realInput.lines())
        assertEquals(42588603, calculateRunningFaster(res.first, res.second))
        assertEquals(42588603, calculateRunningStreamTest(res.first, res.second))
    }


    private fun calculateFaster(duration: Long, record: Long): List<Long> =
        (1..duration).filter { challenge -> isFasterThanRecord(duration, challenge, record) }

    private fun calculateRunningStreamTest(duration: Long, record: Long): Long {
        return (1..duration).fold(0L) { acc, challenge ->
            acc + if (isFasterThanRecord(duration, challenge, record)) 1 else 0
        }
    }

    private fun calculateRunningFaster(duration: Long, record: Long): Long {
        val counter = AtomicLong()
        for (challenge in 1..duration) {
            if (isFasterThanRecord(duration, challenge, record)) {
                counter.incrementAndGet()
            }
        }
        return counter.get()
    }

    private fun isFasterThanRecord(duration: Long, challenge: Long, record: Long) =
        ((duration - challenge) * challenge) > record

    private fun parse(input: List<String>): List<Pair<Long, Long>> =
        input.map { it.split(Pattern.compile("\\s+")).drop(1) }.map { strings -> strings.map { it.toLong() } }.let {
            it[0].zip(it[1])
        }

    private fun parseSingleRace(input: List<String>): Pair<Long, Long> =
        input.map { it.split(Pattern.compile("\\s+")).drop(1) }.map { strings -> strings.joinToString("").toLong() }
            .let {
                it[0] to it[1]
            }
}