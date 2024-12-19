package day18

import readInput
import java.util.PriorityQueue

const val HEIGHT = 71
const val WIDTH = 71

data class Node(
    val position: Pair<Int, Int>,
    val weight: Int
)

fun findPath(input: List<String>): Boolean {
    val corruptedTiles = input.map { line ->
        val (x, y) = line.split(",").map { it.toInt() }
        Pair(x, y)
    }.toSet()

    val allTiles = (0 until HEIGHT).flatMap { x ->
        (0 until WIDTH).map { y ->
            Pair(x, y)
        }
    }.toSet()
    val validTiles = allTiles.subtract(corruptedTiles)

    val end = Pair(HEIGHT - 1, WIDTH - 1)
    val queue = PriorityQueue<Node>(compareBy { it.weight })
    queue.add(Node(Pair(0, 0), 0))
    val visited = mutableSetOf<Pair<Int, Int>>()
    while (queue.isNotEmpty()) {
        val current = queue.poll()
        if (current.position == end) {
            return true
        }
        if (visited.contains(current.position)) continue
        visited.add(current.position)

        val nextPositions = listOf(
            current.position.first to current.position.second - 1,
            current.position.first + 1 to current.position.second,
            current.position.first to current.position.second + 1,
            current.position.first - 1 to current.position.second
        )

        nextPositions.forEach { nextPosition ->
            if (validTiles.contains(nextPosition) && !visited.contains(nextPosition)) {
                queue.add(Node(nextPosition, current.weight + 1))
            }
        }
    }
    return false
}

fun main() {

    val input = readInput(18)
    val lastTrue = input.indices
        .takeWhile { findPath(input.take(it)) }
        .lastOrNull()

    lastTrue?.let {
        println(input[it])
    }
}