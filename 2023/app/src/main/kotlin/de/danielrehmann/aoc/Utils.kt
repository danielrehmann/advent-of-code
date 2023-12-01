package de.danielrehmann.aoc

class Utils {
    companion object {
        fun readLinesFromFile(path: String) =
            (Utils::class.java.getResource(path) ?: error("could not read file $path")).readText().lines()
    }
}