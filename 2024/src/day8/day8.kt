package day8

import readInput

fun main() {
    val input = readInput(8)

    val groups = mutableMapOf<Char, MutableList<Pair<Int, Int>>>()
    input.forEachIndexed { row, line ->
        line.forEachIndexed { col, char ->
            if (char != '.') {
                groups.getOrPut(char) { mutableListOf() }
                    .add(Pair(row, col))
            }
        }
    }

    fun inBounds(pos: Pair<Int, Int>) =
        pos.first in input.indices && pos.second in 0 until input[0].length

    val antinodes = mutableSetOf<Pair<Int, Int>>()
    groups.forEach { (_, positions) ->
        for (i in positions.indices) {
            for (j in i + 1 until positions.size) {
                val yDiff = positions[i].first - positions[j].first
                val xDiff = positions[i].second - positions[j].second

                fun addAntinodes(direction: Int) {
                    var multiplier = direction
                    var antinode = Pair(
                        positions[i].first + yDiff * multiplier,
                        positions[i].second + xDiff * multiplier
                    )

                    while (inBounds(antinode)) {
                        antinodes.add(antinode)
                        if (direction > 0) multiplier++
                        if (direction < 0) multiplier--
                        antinode = Pair(
                            positions[i].first + yDiff * multiplier,
                            positions[i].second + xDiff * multiplier
                        )
                    }
                }

                antinodes.add(positions[i])
                addAntinodes(-1)
                addAntinodes(1)
            }
        }
    }
    println(antinodes.size)
}