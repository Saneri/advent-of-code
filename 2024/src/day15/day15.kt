package day15

import readInputText
import toCoordinates

data class Map(var grid: MutableMap<Pair<Int, Int>, Char>, var position: Pair<Int, Int>) {}

fun findSpace(
    map: Map,
    command: Char,
    searchPosition: Pair<Int, Int>,
): Boolean {
    val (grid) = map
    val nextPosition = when (command) {
        '>' -> searchPosition.first to searchPosition.second + 1
        'v' -> searchPosition.first + 1 to searchPosition.second
        '<' -> searchPosition.first to searchPosition.second - 1
        '^' -> searchPosition.first - 1 to searchPosition.second
        else -> throw Exception("unknown command: $command")
    }

    when (grid[nextPosition]) {
        '.' -> {
            grid[nextPosition] = grid[searchPosition] ?: throw Exception("null value")
            grid[searchPosition] = '.'
            return true
        }

        '#' -> return false
        'O' -> {
            val found = findSpace(map, command, nextPosition)
            if (found) {
                grid[nextPosition] = grid[searchPosition] ?: throw Exception("null value")
                grid[searchPosition] = '.'
            }
            return found
        }

        else -> throw Exception("unknown char ${grid[nextPosition]}")
    }
}

fun main() {
    val input = readInputText(15).split("\n\n")
    val grid = toCoordinates(input[0].split("\n")).toMutableMap()
    val commands = input[1].replace("\n", "")
    val position = grid.entries.first { it.value == '@' }.key

    val map = Map(grid, position)

    for (command in commands) {
        val found = findSpace(map, command, map.position)
        if (found) {
            map.position = when (command) {
                '>' -> map.position.first to map.position.second + 1
                'v' -> map.position.first + 1 to map.position.second
                '<' -> map.position.first to map.position.second - 1
                '^' -> map.position.first - 1 to map.position.second
                else -> throw Exception("unknown command: $command")
            }
        }
    }
    val sum = map.grid.entries.filter { it.value == 'O' }
        .sumOf { (key, _) -> 100 * key.first + key.second }
    println(sum)

}

fun printMap(map: MutableMap<Pair<Int, Int>, Char>) {
    val maxRow = map.keys.maxOf { it.first }
    val maxCol = map.keys.maxOf { it.second }

    for (row in 0..maxRow) {
        for (col in 0..maxCol) {
            print(map[Pair(row, col)] ?: ' ')
        }
        println()
    }
    println()
}