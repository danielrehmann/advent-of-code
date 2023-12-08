package de.danielrehmann.aoc.day8

import de.danielrehmann.aoc.Utils
import org.junit.jupiter.api.Test
import java.util.regex.Pattern
import kotlin.test.assertEquals

class EightTest {
    @Test
    fun canParseMap() {
        val given = """
            RL

            AAA = (BBB, CCC)
            BBB = (DDD, EEE)
            CCC = (ZZZ, GGG)
            DDD = (DDD, DDD)
            EEE = (EEE, EEE)
            GGG = (GGG, GGG)
            ZZZ = (ZZZ, ZZZ)
        """.trimIndent()
        val map = given.toDesertMap()

        assertEquals(DesertMap(listOf(Instruction.Right, Instruction.Left),
                listOf(
                        Node(id = "AAA", leftNode = "BBB", rightNode = "CCC"),
                        Node(id = "BBB", leftNode = "DDD", rightNode = "EEE"),
                        Node(id = "CCC", leftNode = "ZZZ", rightNode = "GGG"),
                        Node(id = "DDD", leftNode = "DDD", rightNode = "DDD"),
                        Node(id = "EEE", leftNode = "EEE", rightNode = "EEE"),
                        Node(id = "GGG", leftNode = "GGG", rightNode = "GGG"),
                        Node(id = "ZZZ", leftNode = "ZZZ", rightNode = "ZZZ")
                )), map)
        assertEquals(2, map.countStepsToZZZ())
    }

    @Test
    fun canCountRepeated() {
        val given = """
            LLR

            AAA = (BBB, BBB)
            BBB = (AAA, ZZZ)
            ZZZ = (ZZZ, ZZZ)
        """.trimIndent()
        val map = given.toDesertMap()
        assertEquals(6, map.countStepsToZZZ())
    }

    @Test
    fun resultPart1() {
        val map = Utils.readFullFile("/8/input1.txt").toDesertMap()
        assertEquals(21883, map.countStepsToZZZ())
    }

    //part 2
    @Test
    fun canNavigateGhostly() {
        val given = """
            LR

            11A = (11B, XXX)
            11B = (XXX, 11Z)
            11Z = (11B, XXX)
            22A = (22B, XXX)
            22B = (22C, 22C)
            22C = (22Z, 22Z)
            22Z = (22B, 22B)
            XXX = (XXX, XXX)
        """.trimIndent()
        val map = given.toDesertMap()
        assertEquals(6, map.countStepsToZGhostlySlow())
    }

    @Test
    fun resultNavigateGhostly() {
        val map = Utils.readFullFile("/8/input1.txt").toDesertMap()
        val steps = map.countStepsToZGhostlyFastRequiresSameDistance()
        assertEquals(12833235391111, steps)
    }

}

data class DesertMap(val instructions: List<Instruction>, val nodes: List<Node>) {
    private val nodeMap = nodes.associate { it.id to (it.leftNode to it.rightNode) }

    fun countStepsToZZZ(): Int {
        var current = "AAA"
        val found = generateSequence { instructions }.flatten().mapIndexed { index, instruction ->
            if (current == "ZZZ") {
                Res.Found(index)
            } else {
                current = instruction.next(nodeMap[current]!!)
                Res.NotFound
            }
        }.filterIsInstance<Res.Found>().take(1).single()
        return found.steps
    }

    fun countStepsToZGhostlyFastRequiresSameDistance(): Long {
        val initial = nodeMap.keys.filter { it.last() == 'A' }
        val listOfFound = initial.map {
            generateSequence { instructions }.flatten().runningFold(it) { acc, instruction ->
                instruction.next(nodeMap[acc]!!)
            }.mapIndexed { index, s ->
                if (s.last() == 'Z') {
                    Res.Found(index)
                } else {
                    Res.NotFound
                }
            }.filterIsInstance<Res.Found>().map { found -> found.steps }.take(1).single()
        }
        val min = listOfFound.min().toLong()
        return (min..Long.MAX_VALUE step min).first { l: Long -> listOfFound.all { l % it == 0L } }
    }


    fun countStepsToZGhostlySlow(): Int {
        val begin = nodeMap.keys.filter { it.last() == 'A' }
        val found = generateSequence { instructions }.flatten()
                .runningFold(begin) { acc, instruction -> acc.map { instruction.next(nodeMap[it]!!) } }
                .mapIndexed { index, strings ->
                    if (strings.all { it.last() == 'Z' }) {
                        Res.Found(index)
                    } else {
                        Res.NotFound
                    }
                }.filterIsInstance<Res.Found>().take(1).single()
        return found.steps
    }

    sealed interface Res {
        data object NotFound : Res
        data class Found(val steps: Int) : Res
    }

    companion object {
        fun of(instructionsString: String, nodeStrings: List<String>): DesertMap {
            val instructions = instructionsString.map {
                when (it) {
                    'R' -> Instruction.Right
                    'L' -> Instruction.Left
                    else -> error("Unknown instruction $it")
                }
            }
            val nodes = nodeStrings.mapNotNull(nodeRegex::matchEntire).map { Node(it.groupValues[1], it.groupValues[2], it.groupValues[3]) }
            return DesertMap(instructions, nodes)
        }

        private val nodeRegex = Regex("^([A-Z0-9]{3}) = \\(([A-Z0-9]{3}), ([A-Z0-9]{3})\\)\$")
    }
}

fun String.toDesertMap(): DesertMap {
    val (instruction, nodesStrings) = split(Pattern.compile("\\n\\n")).let { (instruction, nodesString) -> instruction to nodesString.lines() }
    return DesertMap.of(instruction, nodesStrings)
}

sealed interface Instruction {
    fun next(pair: Pair<String, String>): String

    data object Left : Instruction {
        override fun next(pair: Pair<String, String>): String {
            return pair.first
        }
    }

    data object Right : Instruction {
        override fun next(pair: Pair<String, String>): String {
            return pair.second
        }
    }
}

data class Node(val id: String, val leftNode: String, val rightNode: String)
