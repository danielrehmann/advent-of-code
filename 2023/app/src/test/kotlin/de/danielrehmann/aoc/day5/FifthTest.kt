package de.danielrehmann.aoc.day5

import de.danielrehmann.aoc.Utils
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.async
import kotlinx.coroutines.awaitAll
import kotlinx.coroutines.flow.asFlow
import kotlinx.coroutines.runBlocking
import org.junit.jupiter.api.Test
import kotlin.math.min
import kotlin.test.assertEquals

class FifthTest {
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

    data class Range(val destination: Long, val source: Long, val range: Long) {
        private val before = 0..<source
        private val affected = source..<source + range
        private val after = source + range..Long.MAX_VALUE
        fun findSpot(pos: Long): Long? {
            return when (pos) {
                in affected -> pos - (source - destination)
                else -> null
            }
        }

        fun find(input: LongRange): Pair<MutableList<LongRange>, MutableList<LongRange>> {
            val found = mutableListOf<LongRange>()
            val missed = mutableListOf<LongRange>()

            if (input.first < source) {
                missed.add(input.first..minOf(source - 1, input.last))
            }
            if (input.first >= source && input.first < source + range) {
                val affected = maxOf(input.first, source)..minOf(input.last, source + range - 1)
                found.add(affected.first - (source - destination)..affected.last - (source - destination))
            }
            if (input.last > source + range + 1) {
                missed.add(maxOf(input.first, source + range + 1)..input.last)
            }
            return missed to found
        }
    }

    data class WorkOrder(val name: String, val ranges: List<Range>) {
        fun findSpot(pos: Long): Long {
            return ranges.mapNotNull {
                it.findSpot(pos)
            }.singleOrNull() ?: pos
        }


        companion object {
            fun of(strings: List<String>): WorkOrder {
                return WorkOrder(
                    strings.first(),
                    strings.subList(1, strings.size).let { stringList ->
                        stringList.map { s ->
                            val (destinationRange, sourceRange, rangeLength) = s.split(" ").map { it.toLong() }
                            Range(destinationRange, sourceRange, rangeLength)
                        }
                    })
            }
        }
    }

    data class Fold(val current: List<List<String>>, val newElement: Boolean = true) {
        val seeds: List<Long>
            get() = current.first().first()
                .let { seedRegex.matchEntire(it)!!.groupValues[1].split(" ").map { s -> s.toLong() } }
        val workOrders: List<WorkOrder>
            get() = current.subList(1, current.size).map {
                WorkOrder.of(it)
            }

        fun compile(): Pair<List<Long>, List<WorkOrder>> = seeds to workOrders
        fun add(string: String): Fold {
            return if (string.isBlank()) {
                Fold(current, true)
            } else {
                current.toMutableList().let {
                    if (newElement) {
                        it.add(listOf(string))
                    } else {
                        it[it.lastIndex] = it.last().plus(string)
                    }
                    Fold(it, false)
                }
            }
        }


        companion object {
            private val seedRegex = Regex("^seeds: ([\\d|\\s]+)\$")
        }
    }

    @Test
    fun listOfLinesCanBeSplitInLogicalGroups() {
        val fold = input.lines().fold(Fold(emptyList())) { acc, s -> acc.add(s) }
        assertEquals(listOf<Long>(79, 14, 55, 13), fold.seeds)
        assertEquals(
            WorkOrder(
                name = "seed-to-soil map:",
                ranges = listOf(
                    Range(destination = 50, source = 98, range = 2),
                    Range(destination = 52, source = 50, range = 48)
                )
            ), fold.workOrders.first()
        )

        assertEquals(
            WorkOrder(
                name = "soil-to-fertilizer map:",
                ranges = listOf(
                    Range(destination = 0, source = 15, range = 37),
                    Range(destination = 37, source = 52, range = 2),
                    Range(destination = 39, source = 0, range = 15)
                )
            ), fold.workOrders[1]
        )
    }

    @Test
    fun workOrderCalculatesNextSpot() {
        val order = WorkOrder(
            name = "seed-to-soil map:",
            ranges = listOf(
                Range(destination = 50, source = 98, range = 2),
                Range(destination = 52, source = 50, range = 48)
            )
        )

        assertEquals(81, order.findSpot(79))
        assertEquals(14, order.findSpot(14))
        assertEquals(57, order.findSpot(55))
        assertEquals(13, order.findSpot(13))
    }

    @Test
    fun workOrderCalculatesFullSet() {
        val (seeds, workOrders) = input.lines().fold(Fold(emptyList())) { acc, s -> acc.add(s) }.compile()

        val seedSpots = seeds.map { workOrders.runningFold(it) { acc, workOrder -> workOrder.findSpot(acc) } }
        assertEquals(listOf<Long>(79, 81, 81, 81, 74, 78, 78, 82), seedSpots.first())

        assertEquals(35L, seedSpots.minOfOrNull { it.last() })
    }

    @Test
    fun workOrderCalculatesResult() {
        val (seeds, workOrders) = Utils.readLinesFromFile("/5/input1.txt")
            .fold(Fold(emptyList())) { acc, s -> acc.add(s) }.compile()

        val seedSpots = seeds.map { workOrders.runningFold(it) { acc, workOrder -> workOrder.findSpot(acc) } }
        assertEquals(35L, seedSpots.minOfOrNull { it.last() })
    }

    // part 2
    @Test
    fun workWithSeedRanges() {
        val (seeds, workOrders) = input.lines()
            .fold(Fold(emptyList())) { acc, s -> acc.add(s) }
            .compile()

        val seedRanges = seeds.chunked(2).flatMap { it.first()..<(it.sum()) }
        val seedSpots = seedRanges.map { workOrders.runningFold(it) { acc, workOrder -> workOrder.findSpot(acc) } }
        assertEquals(46L, seedSpots.minOfOrNull { it.last() })

        val seedRanges2 = seeds.windowed(2, 2, false) {
            it.first()..<it.sum()
        }
        assertEquals(listOf(79L..92, 55L..67), seedRanges2)


    }

    @Test
    fun singleRangeSeed() {
        val order = WorkOrder(
            name = "seed-to-soil map:",
            ranges = listOf(
                Range(destination = 50, source = 98, range = 2),
                Range(destination = 52, source = 50, range = 48)
            )
        )
    }

    @Test
    fun workWithSeedRangesResult() {
        val (seeds, workOrders) = Utils.readLinesFromFile("/5/input1.txt")
            .fold(Fold(emptyList())) { acc, s -> acc.add(s) }.compile()
        val seedRanges2 = seeds.windowed(2, 2, false) {
            it.first()..<it.sum()
        }
        assertEquals(768975L..37650595, seedRanges2.first())
        assertEquals(36881621, (768975L..37650595).count())
        val knownSeeds = mutableMapOf<Long, Long>()

        (768975L..37650595).minOf {
            workOrders.fold(it) { acc, workOrder -> workOrder.findSpot(acc) }
        }

//        val knownSeeds = mutableMapOf<Long, Long>()
//        val seedRanges = seeds.chunked(2)
//            .minOfOrNull { longs ->
//                (longs.first()..<(longs.sum())).minOf {
//                    knownSeeds.computeIfAbsent(it) { key ->
//                        workOrders.fold(key) { acc, workOrder -> workOrder.findSpot(acc) }
//                    }
//                }
//            }
//        assertEquals(46L, seedRanges)
    }
}