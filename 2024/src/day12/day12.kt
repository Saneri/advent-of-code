package day12

import readInput

fun main() {
    val map: Map<Pair<Int, Int>, Char> = readInput(12)
        .flatMapIndexed { rowIndex, row ->
            row.mapIndexed { colIndex, char ->
                Pair(rowIndex, colIndex) to char
            }
        }
        .toMap()

    var price = 0
    val visited = mutableSetOf<Pair<Int, Int>>()

    fun iter(node: Pair<Int, Int>, type: Char): Pair<Int, Int> {
        if (visited.contains(node)) return Pair(0, 0)

        visited.add(node)
        var borders = 0
        var area = 1

        val directions = listOf(
            Pair(node.first + 1, node.second),
            Pair(node.first, node.second + 1),
            Pair(node.first - 1, node.second),
            Pair(node.first, node.second - 1),
        )

        for (direction in directions) {
            if (map[direction] == type) {
                val (b, a) = iter(direction, type)
                borders += b
                area += a
            } else {
                borders++
            }
        }

        return Pair(borders, area)
    }

    for (node in map.keys) {
        if (visited.contains(node)) continue
        val (borders, area) = iter(node, map[node]!!)
        price += borders * area
    }
    println(price)
}