package day10

import readInput

fun main() {
    val map = readInput(10).flatMapIndexed { rowIndex, row ->
        row.mapIndexed { colIndex, char ->
            Pair(rowIndex, colIndex) to char
        }
    }.toMap()

    fun iter(position: Pair<Int, Int>, level: Int): Int {
        if (level > 9) {
            return 1
        }
        var trails = 0
        val up = Pair(position.first - 1, position.second)
        if (map[up] == level.digitToChar()) trails += iter(up, level + 1)
        val right = Pair(position.first, position.second + 1)
        if (map[right] == level.digitToChar()) trails += iter(right, level + 1)
        val down = Pair(position.first + 1, position.second)
        if (map[down] == level.digitToChar()) trails += iter(down, level + 1)
        val left = Pair(position.first, position.second - 1)
        if (map[left] == level.digitToChar()) trails += iter(left, level + 1)
        return trails
    }

    val trailheads = map.filter { it.value == '0' }.keys.toList()
    var count = 0
    for (trailhead in trailheads) {
        count += iter(trailhead, 1)
    }
    println(count)
}