package de.danielrehmann.aoc.day2

import de.danielrehmann.aoc.Utils
import org.junit.jupiter.api.Test
import kotlin.test.assertEquals
import kotlin.test.assertFalse


class SecondOfDecemberTest {
    private val game1 = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
    private val game3 = "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\n"

    val fullInput = """
        Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
        Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
        Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
        Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
        Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
    """.trimIndent()

    val gameLimitTest = Cubes(red = 12, green = 13, blue = 14)

    @Test
    fun singleLineCanBeParsed() {
        val gameResult = Game.of(game1)
        assertEquals(
            Game(
                1,
                listOf(Cubes(blue = 3, red = 4), Cubes(red = 1, green = 2, blue = 6), Cubes(green = 2))
            ), gameResult
        )
        assertEquals(true, gameLimitTest.validGame(gameResult))
    }


    @Test
    fun invalidGameIsDetected() {
        val given = game3
        val gameResult = Game.of(given)
        assertFalse(gameLimitTest.validGame(gameResult))
    }

    @Test
    fun fullInputWorks() {
        val result =
            fullInput.lines().map { Game.of(it) }.filter { game -> gameLimitTest.validGame(game) }.sumOf { it.id }

        assertEquals(8, result)
    }

    @Test
    fun fullGame() {
        val result = Utils.readLinesFromFile("/2/input1.txt").map { Game.of(it) }
            .filter { game -> gameLimitTest.validGame(game) }.sumOf { it.id }

        assertEquals(3035, result)
    }

    @Test
    fun gameOneMinRequiredCubes() {
        val game = Game.of(game1)
        assertEquals(Cubes(4, 2, 6), game.minimalRequiredCubes)
        assertEquals(48, game.minimalRequiredCubes.power())
    }

    @Test
    fun testInputPart2() {
        assertEquals(2286, fullInput.lines().map { Game.of(it).minimalRequiredCubes }.sumOf { it.power() })
    }

    @Test
    fun solutionPart2() {
        assertEquals(66027,
            Utils.readLinesFromFile("/2/input1.txt")
                .map { Game.of(it).minimalRequiredCubes }
                .sumOf { it.power() })
    }

}

data class Cubes(val red: Int = 0, val green: Int = 0, val blue: Int = 0) {
    fun validGame(game: Game): Boolean {
        return game.reveals.any {
            it.red > red || it.blue > blue || it.green > green
        }.not()
    }

    fun power() = red * green * blue

    companion object {
        fun of(strings: List<String>): Cubes {
            return strings.associate {
                val (left, right) = it.split(" ")
                right to left.toInt()
            }.let {
                Cubes(
                    blue = it.getOrDefault("blue", 0),
                    red = it.getOrDefault("red", 0),
                    green = it.getOrDefault("green", 0)
                )
            }
        }
    }
}

data class Game(val id: Int, val reveals: List<Cubes>) {
    val minimalRequiredCubes: Cubes
        get() {
            val green = reveals.maxOf { it.green }
            val red = reveals.maxOf { it.red }
            val blue = reveals.maxOf { it.blue }
            return Cubes(red, green, blue)
        }

    companion object {
        val gameRegex = Regex("^Game (\\d+): ([0-9a-z ,;]+)\n?\$")
        fun of(fullString: String): Game {
            val matchEntire = gameRegex.matchEntire(fullString)
            val gameNumber = matchEntire!!.groupValues[1].toInt()
            val reveals = matchEntire.groupValues[2]
            return of(gameNumber, reveals)
        }

        fun of(gameId: Int, reveals: String): Game {
            return Game(gameId, reveals.split("; ").map { it.split(", ") }
                .map {
                    Cubes.of(it)
                })
        }
    }
}
