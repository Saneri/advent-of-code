package day12

import readInput
import toCoordinates

fun main() {
    val map = toCoordinates(readInput(12))

    var price = 0
    val visited = mutableSetOf<Pair<Int, Int>>()

    fun iter(node: Pair<Int, Int>, type: Char): Pair<Int, Int> {
        if (visited.contains(node)) return Pair(0, 0)

        visited.add(node)
        var corners = 0
        var area = 1

        val directions = listOf(
            Pair(node.first + 1, node.second),
            Pair(node.first, node.second + 1),
            Pair(node.first - 1, node.second),
            Pair(node.first, node.second - 1),
        )

        val adjacentPairs = directions.indices.map { i ->
            Pair(directions[i], directions[(i + 1) % directions.size])
        }

        for ((dir1, dir2) in adjacentPairs) {
            // convex corners
            if (map[dir1] != type && map[dir2] != type) {
                corners++
            }
            // concave corners
            val diagonal = Pair(
                dir1.first + dir2.first - node.first,
                dir1.second + dir2.second - node.second
            )
            if (map[dir1] == type && map[dir2] == type && map[diagonal] != type) {
                corners++
            }
        }

        for (direction in directions) {
            if (map[direction] == type) {
                val (b, a) = iter(direction, type)
                corners += b
                area += a
            }
        }

        return Pair(corners, area)
    }

    for (node in map.keys) {
        if (visited.contains(node)) continue
        val (corners, area) = iter(node, map[node]!!)
        price += corners * area
    }
    println(price)
}